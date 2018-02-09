/*
* SensorWalkSkill implements a skill
* that makes the HEXA to walk forward and change direction
* when encountering obstacles
 */
package examples

import (
	"math"
	"os"
	"time"

	"mind/core/framework/drivers/distance"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	TIME_TO_NEXT_REACTION = 2000 // milliseconds
	DISTANCE_TO_REACTION  = 250  // millimeters
	MOVE_HEAD_DURATION    = 500  // milliseconds
	ROTATE_DEGREES        = 130  // degrees out of 360
	WALK_SPEED            = 1.0  // cm per second
	SENSE_INTERVAL        = 250  // four times per second
)

func newDirection(direction float64) float64 {
	return math.Mod(direction+ROTATE_DEGREES, 360) * -1
}

type SensorWalkSkill struct {
	skill.Base
	stop      chan bool
	direction float64
}

func NewSkill() skill.Interface {
	return &SensorWalkSkill{
		stop: make(chan bool),
	}
}

func (d *SensorWalkSkill) distance() float64 {
	distance, err := distance.Value()
	if err != nil {
		log.Error.Println(err)
	}
	return distance
}

func (d *SensorWalkSkill) changeDirection() {
	d.direction = newDirection(d.direction)
	hexabody.MoveHead(d.direction, MOVE_HEAD_DURATION)
	hexabody.WalkContinuously(0, WALK_SPEED)
	time.Sleep(TIME_TO_NEXT_REACTION * time.Millisecond)
}

func (d *SensorWalkSkill) shouldChangeDirection() bool {
	return d.distance() < DISTANCE_TO_REACTION
}

func (d *SensorWalkSkill) walk() {
	hexabody.WalkContinuously(0, WALK_SPEED)
	for {
		select {
		case <-d.stop:
			return
		default:
			if d.shouldChangeDirection() {
				d.changeDirection()
			}
			time.Sleep(SENSE_INTERVAL * time.Millisecond)
		}
	}
}

func (d *SensorWalkSkill) OnStart() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}
	err = distance.Start()
	if err != nil {
		log.Error.Println("Distance start err:", err)
		return
	}
	if !distance.Available() {
		log.Error.Println("Distance sensor is not available")
	}
}

func (d *SensorWalkSkill) OnClose() {
	hexabody.Close()
	distance.Close()
}

func (d *SensorWalkSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *SensorWalkSkill) OnRecvString(data string) {
	switch data {
	case "start":
		go d.walk()
	case "stop":
		d.stop <- true
		hexabody.StopWalkingContinuously()
		hexabody.Relax()
	}
}
