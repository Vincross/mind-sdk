package AudioSkill

import (
	"mind/core/framework/drivers/audio"
	"mind/core/framework/log"
	"mind/core/framework/skill"

	"bufio"
	"io"
	"os"
	"path/filepath"
	"sync"
)

const (
	RecordFileName = "record.raw"
	OriginFilePath = "deps/testmusic.wav"

	StatusStop               = "stop"
	StatusListenAndPlay      = "listenAndPlay"
	StatusRecordToFile       = "recordToFile"
	StatusPlayWithFile       = "playWithFile"
	StatusPlayWithOriginFile = "playWithOriginFile"
)

type AudioSkill struct {
	skill.Base
	stopRecordChan        chan bool
	stopListenAndPlayChan chan bool
	stopPlayChan          chan bool
	workStatus            string
	statusMutex           sync.Mutex
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.
	return &AudioSkill{
		stopRecordChan:        make(chan bool, 1),
		stopListenAndPlayChan: make(chan bool, 1),
		stopPlayChan:          make(chan bool, 1),
		workStatus:            StatusStop,
	}
}

func (d *AudioSkill) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *AudioSkill) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d *AudioSkill) OnConnect() {
	// Use this method to do something when the remote connected.
	log.Info.Println("connect to skill...")
}

func (d *AudioSkill) RecordToFile() {
	// Get the directory on HEXA of your skill. It will be removed and built every time you run 'mind run'.
	log.Info.Println("Start record...")
	defer log.Info.Println("Finish record...")
	dataPath, err := skill.SkillDataPath()
	if err != nil {
		log.Error.Println("Get data path error:", err)
		return
	}
	log.Info.Println("Get data path:", dataPath)

	err = audio.Start()
	if err != nil {
		log.Error.Println("Start audio driver error:", err)
		return
	}
	defer audio.Close()
	audio.Init(1, 44100, audio.FormatS16LE)

	os.Remove(filepath.Join(dataPath, RecordFileName))
	file, err := os.OpenFile(filepath.Join(dataPath, RecordFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Error.Println("Open file error:", err)
		return
	}
	defer file.Close()
	total := []byte{}
Loop:
	for {
		select {
		case <-d.stopRecordChan:
			break Loop
		default:
			readBuf, err := audio.Read()
			if err != nil {
				log.Error.Println("Audio read error:", err)
			}
			total = append(total, readBuf...)
		}
	}
	file.Write(total)
}

func (d *AudioSkill) PlayWithFile(filepath string) {
	log.Info.Println("Play file:", filepath)
	defer log.Info.Println("Finish play:", filepath)
	fi, err := os.Open(filepath)
	if err != nil {
		log.Error.Println("open filepath error:", err)
		return
	}
	err = audio.Start()
	if err != nil {
		log.Error.Println("Start audio driver error:", err)
	}
	defer audio.Close()
	audio.Init(1, 44100, audio.FormatS16LE)
	buf := make([]byte, 1024)
	r := bufio.NewReader(fi)
Loop:
	for {
		select {
		case <-d.stopPlayChan:
			break Loop
		default:
			_, err := r.Read(buf)
			if err != nil {
				if err == io.EOF {
					break Loop
				} else {
					log.Error.Println("read err:", err)
					return
				}
			}
			audio.Write(buf)
		}
	}

}

func (d *AudioSkill) ListenAndPlay() {
	log.Info.Println("Start Record and Play loop")
	defer log.Info.Println("Finish Record and Play loop")
	err := audio.Start()
	if err != nil {
		log.Error.Println("Start audio driver error:", err)
	}
	defer audio.Close()
	audio.Init(1, 44100, audio.FormatS16LE)
	readBuf := make([]byte, 1024)
Loop:
	for {
		select {
		case <-d.stopListenAndPlayChan:
			break Loop
		default:
			readBuf, err = audio.Read()
			if err != nil {
				log.Error.Println("read err:", err)
			}
			audio.Write(readBuf)
		}
	}
}

func (d *AudioSkill) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *AudioSkill) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *AudioSkill) Stop() {
	d.statusMutex.Lock()
	defer d.statusMutex.Unlock()

	switch d.workStatus {
	case StatusStop:
	case StatusPlayWithFile:
		fallthrough
	case StatusPlayWithOriginFile:
		d.stopPlayChan <- true
	case StatusRecordToFile:
		d.stopRecordChan <- true
	case StatusListenAndPlay:
		d.stopListenAndPlayChan <- true
	default:
	}
	d.workStatus = StatusStop
}

func (d *AudioSkill) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
	log.Info.Println("Get command:", data)
	if d.workStatus != StatusStop && data != StatusStop {
		log.Warn.Println("Running other work:", d.workStatus)
		return
	}
	switch data {
	case StatusStop:
		d.Stop()
	case StatusListenAndPlay:
		d.workStatus = StatusListenAndPlay
		d.ListenAndPlay()
	case StatusRecordToFile:
		d.workStatus = StatusRecordToFile
		d.RecordToFile()
	case StatusPlayWithFile:
		d.workStatus = StatusPlayWithFile
		dataPath, err := skill.SkillDataPath()
		if err != nil {
			log.Error.Println("Get data path error:", err)
			return
		}
		d.PlayWithFile(filepath.Join(dataPath, RecordFileName))
	case StatusPlayWithOriginFile:
		d.workStatus = StatusPlayWithOriginFile
		d.PlayWithFile(OriginFilePath)
	default:
	}
	d.workStatus = StatusStop
}
