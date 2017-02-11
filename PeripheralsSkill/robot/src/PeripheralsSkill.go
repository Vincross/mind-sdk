/*
* PeripheralsSkill implements a skill that shows how to use peripherals interface.
* This example use pca9685 as i2c device, i2c part of it need to be modified according to specific device.
 */
package example

import (
	"mind/core/framework/drivers/adc"
	"mind/core/framework/drivers/gpio"
	"mind/core/framework/drivers/i2c"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

type PeripheralsSkill struct {
	skill.Base
}

func NewSkill() skill.Interface {
	return &PeripheralsSkill{}
}

func runADC() {
	adc.Start()
	for i := 0; i < 4; i++ {
		log.Info.Println(adc.Value(i))
	}
	adc.Close()
}

func runGPIO() {
	gpio.Start()
	for i := 0; i < 4; i++ {
		log.Info.Println("GPIO PIN:", i)
		gpio.Output(i, i%2 == 0)
		log.Info.Println(gpio.High(i))
	}
	gpio.Close()
}

func runI2C() {
	i2c.Start()
	i2c.Set(0x41, 0x00, 0x00)
	log.Info.Println(i2c.Value(0x41, 0x00, 1))
	i2c.Set(0x41, 0x14, 0xff)
	log.Info.Println(i2c.Value(0x41, 0x14, 1))
	i2c.Set(0x41, 0x15, 0x0)
	log.Info.Println(i2c.Value(0x41, 0x15, 1))
	i2c.Close()
}

func (d *PeripheralsSkill) OnConnect() {
	runADC()
	runGPIO()
	runI2C()
}
