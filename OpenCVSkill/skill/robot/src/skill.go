package examples

import (
	"mind/core/framework/skill"

	opencv "github.com/lazywei/go-opencv/opencv"
)

type OpenCVSkill struct {
	skill.Base
}

func NewSkill() skill.Interface {
	return &OpenCVSkill{}
}

func (d *OpenCVSkill) OnStart() {
	filename := "assets/bert.jpg"
	srcImg := opencv.LoadImage(filename)
	if srcImg == nil {
		panic("Loading Image failed")
	}
	defer srcImg.Release()
	resized1 := opencv.Resize(srcImg, 400, 0, 0)
	resized2 := opencv.Resize(srcImg, 300, 500, 0)
	resized3 := opencv.Resize(srcImg, 300, 500, 2)
	opencv.SaveImage("resized1.jpg", resized1, 0)
	opencv.SaveImage("resized2.jpg", resized2, 0)
	opencv.SaveImage("resized3.jpg", resized3, 0)
}
