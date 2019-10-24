package demo

import (
	"encoding/json"
	"mind/core/framework/drivers/tank"
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"sync"
)

type demo struct {
	skill.Base
}

var skillMutex sync.RWMutex

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &demo{}
}

func (d *demo) OnStart() {
	// Use this method to do something when this skill is starting.
	err := tank.Start()
	if err != nil {
		log.Error.Println("rover start error:", err)
		return
	}
}

func (d *demo) OnClose() {
	// Use this method to do something when this skill is closing.
	tank.Close()
}

func (d *demo) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d *demo) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *demo) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
	skillMutex.Lock()
	defer skillMutex.Unlock()
	log.Info.Println(string(data))

	var jsonObj map[string]interface{}
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		log.Error.Println(err)
		return
	}

	action := d.TransString(jsonObj, "action")
	log.Debug.Println(action)

	if action != "" {
		switch action {
		case "move":
			tank.MoveCtrl(500, 0)
		case "stop":
			tank.MoveCtrl(0, 0)
		default:
		}
	}
}

func (d *demo) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}

func (d *demo) TransString(jsonStr map[string]interface{}, str string) string {
	result, ok := jsonStr[str].(string)
	if ok == false {
		log.Error.Println(str, jsonStr[str], " can't be transfered to string!")
		return ""
	}
	return result
}
