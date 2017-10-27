package SelectGaitSkill

import (
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"os"
)

const (
	WALK_SPEED = 1.0 // cm per second
)

type SelectGaitSkill struct {
	skill.Base
	isRunning bool
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.
	return &SelectGaitSkill{}
}

func (d *SelectGaitSkill) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *SelectGaitSkill) OnClose() {
	// Use this method to do something when this skill is closing.
	hexabody.Close()
}

func (d *SelectGaitSkill) OnConnect() {
	// Use this method to do something when the remote connected.
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("start hexabody driver error:", err)
		os.Exit(0)
	}
}

func (d *SelectGaitSkill) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *SelectGaitSkill) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *SelectGaitSkill) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
	switch data {
	case "start":
		d.isRunning = true
		hexabody.WalkContinuously(hexabody.Direction(), WALK_SPEED)
	case "stop":
		d.isRunning = false
		hexabody.StopWalkingContinuously()
		hexabody.Relax()
	case "GaitWave":
		d.changeGait(hexabody.GaitWave)
	case "GaitRipple":
		d.changeGait(hexabody.GaitRipple)
	case "GaitTripod":
		d.changeGait(hexabody.GaitTripod)
	case "GaitOriginal":
		d.changeGait(hexabody.GaitOriginal)
	}
}

func (d *SelectGaitSkill) changeGait(gait hexabody.GaitType) {
	if d.isRunning {
		hexabody.StopWalkingContinuously()
	}
	hexabody.SelectGait(gait)
	if d.isRunning {
		hexabody.WalkContinuously(hexabody.Direction(), WALK_SPEED)
	}
}
