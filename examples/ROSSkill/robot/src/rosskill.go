// ROS Skill implements a skill that shows how to publish images to an ROS topic via rosserial.

package example

// #include "rosskill.h"
import "C"

import (
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
	"mind/core/framework/skill"

	"bytes"
	"image/jpeg"
)

const (
	rosMasterIP = "127.0.0.1" // ROS_MASTER_IP need to be modified manually
	rosTopic    = "/image/compressed"
)

var jpegOption = &jpeg.Options{10} // Size of rosserial messages can not exceed 32767, so we have to use low quality images here.

type rosskill struct {
	skill.Base
	imagePublisher *C.ImagePublisher
	stop           chan bool
}

func NewSkill() skill.Interface {
	return &rosskill{
		imagePublisher: C.NewImagePublisher(C.CString(rosMasterIP), C.CString(rosTopic)),
		stop:           make(chan bool),
	}
}

func (d *rosskill) publishImages() {
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		var b bytes.Buffer
		img := media.SnapshotRGBA()
		jpeg.Encode(&b, img, jpegOption)
		C.PublishImage(
			d.imagePublisher,
			(*C.uchar)(C.CBytes(b.Bytes())),
			C.int(len(b.Bytes())),
		)
		log.Info.Println("Sent Image with length:", len(b.Bytes()))
	}
}

func (d *rosskill) OnStart() {
	err := media.Start()
	if err != nil {
		log.Error.Println("Media start err:", err)
		return
	}
	d.publishImages()
}

func (d *rosskill) OnClose() {
	d.stop <- true
	C.DeleteImagePublisher(d.imagePublisher)
	media.Close()
}
