# MIND SDK
![](https://img.shields.io/badge/master-0.4.0-green.svg?style=flat)

This repository contains everything needed to develop *Skills* for the [HEXA](https://www.vincross.com/hexa)

Check out the [Introduction and Getting Started guides](https://www.vincross.com/developer/introduction/mind-overview)

## [MIND Command-line Interface](cli)
Command-line Interface used for *Skill* development.

```
MIND Command-line Interface v0.4.0

Usage:
  mind [command]

Available Commands:
  build             Build a Skill
  get-default-robot Returns the name of the default robot
  init              Initialize and scaffold a new Skill
  login             Authenticate yourself
  pack              Pack a Skill
  run               Run Skill on robot
  scan              Scan your network or specific IP for robots
  set-default-robot Set the default robot name
  upgrade           Upgrades mindcli container to latest version
  x                 Run a command inside of a cross-compiling capable docker container

Use "mind [command] --help" for more information about a command.
```
## [XCompile](xcompile)
A Docker container with everything needed to cross-compile C/C++ and Golang applications to the ARM architecture.

## [Examples](examples)
*Skill* examples to learn from.
