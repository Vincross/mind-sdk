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
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	FAST_DURATION = 80
	SLOW_DURATION = 500
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
	hexabody.MoveHead(0.0, FAST_DURATION)
	// Using goroutines to make some commands be executed at the same time
	legPosition2 := hexabody.NewLegPosition().SetCoordinates(-100, 50.0, 70.0)
	// legPositionInfo() is used to check, adjust or return the legposition's infomation and it's not necessary.
	legPositionInfo(legPosition2)
	// Moveleg
	go hexabody.MoveLeg(2, legPosition2, SLOW_DURATION)
	hexabody.MoveLeg(5, hexabody.NewLegPosition().SetCoordinates(100, 50.0, 70.0), SLOW_DURATION)
	go hexabody.MoveJoint(0, 1, 90, SLOW_DURATION)
	hexabody.MoveJoint(0, 2, 45, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 90, FAST_DURATION)
	hexabody.MoveJoint(1, 2, 45, FAST_DURATION)
}

func legPositionInfo(legPosition *hexabody.LegPosition) {
	if !legPosition.IsValid() {
		log.Info.Println("The position is not valid, means it's unreachale, fit it.")
		legPosition.Fit()
	}
	x, y, z, err := legPosition.Coordinates()
	if err != nil {
		log.Info.Println("Get coordinates of legposition error:", err)
	} else {
		log.Info.Println("The coordinates of legposition are:", x, y, z)
	}
}

func moveLegs(v float64) {
	go hexabody.MoveJoint(0, 1, 45*math.Sin(v*math.Pi/180)+70, FAST_DURATION)
	go hexabody.MoveJoint(0, 0, 35*math.Cos(v*math.Pi/180)+60, FAST_DURATION)
	hexabody.MoveJoint(1, 1, 45*math.Cos(v*math.Pi/180)+70, FAST_DURATION)
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
			v += 10
		}
	}
}
func (d *MoveLegsSkill) OnStart() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("hexabody can't start:", err)
	}
}

func (d *MoveLegsSkill) OnClose() {
	hexabody.Close()
}

func (d *MoveLegsSkill) OnDisconnect() {
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
