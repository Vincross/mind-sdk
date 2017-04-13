package BoilerplateSkillName

import (
	"mind/core/framework/skill"
)

type BoilerplateSkillName struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &BoilerplateSkillName{}
}

func (d *BoilerplateSkillName) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *BoilerplateSkillName) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d *BoilerplateSkillName) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d *BoilerplateSkillName) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *BoilerplateSkillName) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *BoilerplateSkillName) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
