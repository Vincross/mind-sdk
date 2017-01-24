/*
* MoveLegsSkill implements a skill
* that makes the HEXA to move its right
* arm up and down and left arm in a circle
 */

package examples

import (
	"math"
	"os"

	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/skill"
)

type MoveLegsSkill struct {
	skill.Base
	stop chan bool
}

func NewSkill() skill.Interface {
	return &MoveLegsSkill{
		stop: make(chan bool),
	}
}

func ready() {
	hexabody.Stand()
	hexabody.MoveHead(0.0, 1)
	hexabody.MoveLeg(2, hexabody.NewLegPosition(-100, 50.0, 70.0), 1)
	hexabody.MoveLeg(5, hexabody.NewLegPosition(100, 50.0, 70.0), 1)
	hexabody.MoveJoint(0, 1, 90, 1)
	hexabody.MoveJoint(0, 2, 45, 1)
	hexabody.MoveJoint(1, 1, 90, 1)
	hexabody.MoveJoint(1, 2, 45, 1)
}

func moveLegs(v float64) {
	hexabody.MoveJoint(0, 1, 45*math.Sin(v*math.Pi/180)+70, 1)
	hexabody.MoveJoint(0, 0, 45*math.Cos(v*math.Pi/180)+100, 1)
	hexabody.MoveJoint(1, 1, 45*math.Cos(v*math.Pi/180)+70, 1)
}

func (d *MoveLegsSkill) play() {
	ready()
	v := 0.0
	for {
		select {
		case <-d.stop:
			return
		default:
			moveLegs(v)
			v += 1
		}
	}
}
func (d *MoveLegsSkill) OnStart() {
	hexabody.Start()
}

func (d *MoveLegsSkill) OnClose() {
	hexabody.Close()
}

func (d *SightSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *MoveLegsSkill) OnRecvString(data string) {
	switch data {
	case "start":
		go d.play()
	case "stop":
		d.stop <- true
		hexabody.RelaxLegs()
	}
}
