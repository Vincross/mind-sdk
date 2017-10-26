/*
* Sight skill makes the HEXA react to what's infront of it
* by comparing images captured from the camera.
 */
package examples

import (
	"fmt"
	"image"
	"math"
	"os"
	"time"

	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	TIME_BETWEEN_IMAGES = 500
	REACT_ON_DIFF_VALUE = 70000
	STAND_TIME          = 5000
	STAND_DEPTH         = 200
	SIT_DEPTH           = 50
)

type SightSkill struct {
	skill.Base
	stop chan bool
}

func NewSkill() skill.Interface {
	return &SightSkill{
		stop: make(chan bool),
	}
}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}

func FastCompare(img1, img2 *image.RGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("Image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}
	accumError := int64(0)
	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}
	return int64(math.Sqrt(float64(accumError))), nil
}

func (d *SightSkill) sight() {
	lastImage := media.SnapshotRGBA()
	for {
		select {
		case <-d.stop:
			return
		default:
			image := media.SnapshotRGBA()
			diff, err := FastCompare(image, lastImage)
			if err != nil {
				log.Error.Println(err)
			}
			log.Info.Println("Image difference: ", diff)
			if diff >= REACT_ON_DIFF_VALUE {
				hexabody.StandWithHeight(STAND_DEPTH)
				time.Sleep(STAND_TIME * time.Millisecond)
				hexabody.StandWithHeight(SIT_DEPTH)
			}
			time.Sleep(TIME_BETWEEN_IMAGES * 2 * time.Millisecond)
			lastImage = media.SnapshotRGBA()
			time.Sleep(TIME_BETWEEN_IMAGES * time.Millisecond)
		}
	}
}

func (d *SightSkill) OnStart() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}
	if !media.Available() {
		log.Error.Println("Media driver not available")
		return
	}
	if err = media.Start(); err != nil {
		log.Error.Println("Media driver could not start")
	}
}

func (d *SightSkill) OnClose() {
	hexabody.Close()
}

func (d *SightSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *SightSkill) OnRecvString(data string) {
	switch data {
	case "start":
		go d.sight()
	case "stop":
		d.stop <- true
	}
}
