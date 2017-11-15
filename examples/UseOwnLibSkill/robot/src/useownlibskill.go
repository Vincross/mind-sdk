package UseOwnLibSkill

// #include "calculate/calculate.h"
// #cgo LDFLAGS: -lcalculate
import "C"
import (
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

type UseOwnLibSkill struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &UseOwnLibSkill{}
}

func (d *UseOwnLibSkill) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *UseOwnLibSkill) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d *UseOwnLibSkill) OnConnect() {
	// Use this method to do something when the remote connected.
	log.Info.Println(C.addNum(1, 1))
}

func (d *UseOwnLibSkill) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *UseOwnLibSkill) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *UseOwnLibSkill) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
