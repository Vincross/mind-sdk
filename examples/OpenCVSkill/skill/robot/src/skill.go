package examples

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"mind/core/framework"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"os"

	"github.com/lazywei/go-opencv/opencv"
)

type OpenCVSkill struct {
	skill.Base
	stop    chan bool
	cascade *opencv.HaarCascade
}

func NewSkill() skill.Interface {
	return &OpenCVSkill{
		stop:    make(chan bool),
		cascade: opencv.LoadHaarClassifierCascade("assets/haarcascade_frontalface_alt.xml"),
	}
}

func (d *OpenCVSkill) sight() {
	for {
		select {
		case <-d.stop:
			return
		default:
			image := media.SnapshotRGBA()
			buf := new(bytes.Buffer)
			jpeg.Encode(buf, image, nil)
			str := base64.StdEncoding.EncodeToString(buf.Bytes())
			framework.SendString(str)
			cvimg := opencv.FromImageUnsafe(image)
			faces := d.cascade.DetectObjects(cvimg)
			hexabody.StandWithHeight(float64(len(faces)) * 50)
		}
	}
}

func (d *OpenCVSkill) OnStart() {
	log.Info.Println("Started")
}

func (d *OpenCVSkill) OnConnect() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}
	if !media.Available() {
		log.Error.Println("Media driver not available")
		return
	}
	if err := media.Start(); err != nil {
		log.Error.Println("Media driver could not start")
	}
}

func (d *OpenCVSkill) OnClose() {
	hexabody.Close()
}

func (d *OpenCVSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *OpenCVSkill) OnRecvString(data string) {
	log.Info.Println(data)
	switch data {
	case "start":
		go d.sight()
	case "stop":
		d.stop <- true
	}
}
