/*
* BalanceSkill implements a skill
* that makes the HEXA to try to stand horizontally when it is standing on a plane,
* you could also just let it stand on your arms.
*
* It works like a PID controller. see https://en.wikipedia.org/wiki/PID_controller.
 */
package example

import (
	"os"

	"mind/core/framework/drivers/accelerometer"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	PID_KP            = 0.4  // proportional term
	PID_KI            = 0.1  // integral term
	PID_KD            = 0.01 // derivative term
	PID_MIN           = -5.0 // min value of output
	PID_MAX           = 5.0  // max value of output
	EPS               = 3.0  // acceptable error value of angle
	LEG_MOVE_DURATION = 500
)

type BalanceSkill struct {
	skill.Base
	xpid *PID
	ypid *PID
}

func NewSkill() skill.Interface {
	return &BalanceSkill{
		xpid: NewPID(PID_KP, PID_KI, PID_KD, PID_MIN, PID_MAX),
		ypid: NewPID(PID_KP, PID_KI, PID_KD, PID_MIN, PID_MAX),
	}
}

func (b *BalanceSkill) KeepBalance() {
	curx, cury := 0.0, 0.0
	err := accelerometer.Start()
	if err != nil {
		log.Error.Println("Accelerometer start error:", err)
		return
	}
	for {
		_, _, _, x, y, _, err := accelerometer.Value()
		if err != nil {
			log.Error.Println(err)
			return
		}
		log.Debug.Println(x, y)
		if -EPS < x && x < EPS && -EPS < y && y < EPS {
			continue
		}
		curx += b.xpid.Compute(x)
		cury += b.ypid.Compute(y)
		legPositions := hexabody.PitchRoll(curx, cury)
		hexabody.MoveLegs(legPositions, LEG_MOVE_DURATION)
	}
}

func (b *BalanceSkill) OnStart() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("Hexabody start error:", err)
		return
	}

	hexabody.MoveHead(0, 0)
	hexabody.Stand()

	b.KeepBalance()
}

func (b *BalanceSkill) OnDisconnect() {
	os.Exit(0)
}

func (b *BalanceSkill) OnClose() {
	accelerometer.Close()
	hexabody.Close()
}
