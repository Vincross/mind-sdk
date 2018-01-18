# MIND SDK
![](https://img.shields.io/badge/master-0.6.1-green.svg?style=flat)

This repository contains everything needed to develop *Skills* for the [HEXA](https://www.vincross.com/hexa)

Check out the [Introduction and Getting Started guides](https://documentation.vincross.com/Development/installmind.html)


## [MIND Command-line Interface](cli)
Command-line Interface used for *Skill* development.

```
MIND Command-line Interface v0.6.1

Usage:
  mind [command]

Available Commands:
  build                Build a Skill
  flight-test          Flight test a Skill on mobile device
  get-default-robot    Returns the name of the default robot
  get-default-robot-ip Returns the IP of the default robot
  init                 Initialize and scaffold a new Skill
  login                Authenticate yourself
  pack                 Pack a Skill
  run                  Run Skill on robot
  scan                 Scan your network or specific IP for robots
  set-default-robot    Set the default robot name
  upgrade              Upgrades mindcli container to latest version
  x                    Run a command inside of a cross-compiling capable docker container

Use "mind [command] --help" for more information about a command.
```
### Supported platforms
* macOS
* Linux
* Windows
* Windows + Hyper-V

### Build instructions (macOS/Linux)
* Project has to be cloned to somewhere in your `$GOPATH` 
* To install your `$GOBIN` has to be set.
```
$ cd mindsdk/cli
$ make build install 
```

## [XCompile](xcompile)
A Docker image with everything needed to cross-compile C/C++ and Golang applications to the ARM architecture.

## [Examples](examples)
*Skill* examples to learn from.
