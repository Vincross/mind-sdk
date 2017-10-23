package example

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"mind/core/framework"
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"os"
)

type example struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &example{}
}

func (d *example) OnStart() {
}

func (d *example) OnConnect() {
	err := media.Start()
	if err != nil {
		log.Error.Println("Media start err:", err)
		return
	}
	for {
		log.Info.Println("Connected")
		buf := new(bytes.Buffer)
		log.Info.Println("JPEG")
		jpeg.Encode(buf, media.SnapshotYCbCr(), nil)
		log.Info.Println("BASE64")
		str := base64.StdEncoding.EncodeToString(buf.Bytes())
		log.Info.Println("SENDING")
		framework.SendString(str)
		log.Info.Println("Sent:", str[:20], len(str))
	}
}

func (d *example) OnDisconnect() {
	os.Exit(0)
}

func (d *example) OnClose() {
	// Use this method to do something when this skill is closing.
	media.Close()
}

func (d *example) OnRecvJson(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *example) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
