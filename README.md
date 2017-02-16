# Example Skills

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

# Getting started
In order to get started, you’ll need a [HEXA](https://www.vincross.com/hexa).

## Obtaining and configuring the MIND SDK

### Unix-like
Download the MIND CLI to your path.
```
$ sudo curl -o /usr/local/bin/mind https://cdn-static.vincross.com/downloads/mind/mind-`uname -s`-`uname -m`
$ sudo chmod +x /usr/local/bin/mind
```

### Windows
**Download MIND**

Download the [32 bit version of MIND CLI](https://cdn-static.vincross.com/downloads/mind/windows-i386/mind.exe)

Download the [64 bit version of MIND CLI](https://cdn-static.vincross.com/downloads/mind/windows-x86_64/mind.exe)

**Put MIND in your PATH**

To use mind in a convenient way you should put it in your PATH.
We will not go in to how to do this on Windows but Google is your friend :)

### Running the examples
Example of running the SensorWalkSkill example:
```
$ git clone https://github.com/vincross/hexa-example-skills.git
$ cd SensorWalkSkill
$ mind run -r HEXA # Where HEXA is the default name of your robot (this can be changed through the mobile app)
```
The SensorWalkSkill should now be running on your HEXA and display some output in your terminal.
