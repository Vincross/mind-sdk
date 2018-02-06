# Example Skills

## Building and running
In order to build and run these examples you need to [install the MIND SDK](https://www.vincross.com/developer/introduction/getting-started/macos-&-linux)
and configure it for your [HEXA](https://www.vincross.com/hexa).

Example:
```
$ cd SensorWalkSkill
$ mind build
$ mind pack
$ mind run
``` 

## [SensorWalkSkill](SensorWalkSkill)
SensorWalkSkill makes the HEXA to walk forward and change direction when encountering obstacles.

## [MoveLegsSkill](MoveLegsSkill)
MoveLegsSkill makes the HEXA to move its front right leg up and down and front left leg in a circle.

## [BalanceSkill](BalanceSkill)
BalanceSkill makes the HEXA keep its balance when standing on an unstable
surface. You could also try to hold the HEXA's feet with your hands. 

## [SightSkill](SightSkill)
SightSkill makes the HEXA react to what's infront of it by comparing images captured from the camera.

## [PeripheralsSkill](PeripheralsSkill)
PeripheralsSkill implements a skill that shows how to use the peripherals interface.

## [OpenCVSkill](OpenCVSkill)
OpenCVSkill is an example project showing how to cross-compile C++ libraries and bind them to Golang for Skill development. 
This Skill makes the HEXA stand up when it detects human faces.

## [ROSSkill](ROSSkill)
ROSSkill implements a *skill* that shows how to publish images to an ROS topic via rosserial.

## [SelectGaitSkill](SelectGaitSkill)
SelectGaitSkill is an example project showing how to select different gaits for hexa.

## [CameraSkill](CameraSkill)
CameraSkill use the media to show the video on the remote webpage.

## [AudioSkill](AudioSkill)
AudioSkill shows some simple examples that how to use 'audio' driver on HEXA.

