package mindcli

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type MindCliConfig struct {
	Image             string
	ContainerSkillDir string
	ServeMPKPort      int
	ServeRemotePort   int
}

type MindCli struct {
	config       *MindCliConfig
	userConfig   *UserConfig
	RobotScanner *RobotScanner
}

func NewMindCli(robotScanner *RobotScanner, userConfig *UserConfig, config *MindCliConfig) *MindCli {
	if userConfig.DockerImage != "" {
		config.Image = userConfig.DockerImage
	}
	mindcli := &MindCli{
		RobotScanner: robotScanner,
		userConfig:   userConfig,
		config:       config,
	}
	return mindcli
}

func (mindcli *MindCli) skillMountPoint() string {
	curdir, err := os.Getwd()
	if runtime.GOOS == "windows" {
		re := regexp.MustCompile(`^(?P<disk>[A-z]):(?P<path>.*)$`)
		n1 := re.SubexpNames()
		r2 := re.FindAllStringSubmatch(curdir, -1)[0]
		md := map[string]string{}
		for i, n := range r2 {
			md[n1[i]] = n
		}
		return fmt.Sprintf(
			"/%s%s:%s",
			strings.ToLower(md["disk"]),
			strings.Replace(md["path"], "\\", "/", -1),
			mindcli.config.ContainerSkillDir)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return fmt.Sprintf("%s:%s", curdir, mindcli.config.ContainerSkillDir)
}

func (mindcli *MindCli) dockerUpgradeImageArgs() []string {
	return []string{
		"pull",
		fmt.Sprintf("%s", mindcli.config.Image),
	}
}

func mindUser() string {
	if runtime.GOOS == "linux" {
		user, err := user.Current()
		if err != nil {
			return ""
		}
		return user.Uid
	}
	return ""
}

func (mindcli *MindCli) dockerXArgs(args ...string) []string {
	c := []string{
		"run",
		"-it",
		"--rm",
		"-p", fmt.Sprintf("%d:%d", mindcli.config.ServeRemotePort, mindcli.config.ServeRemotePort),
		"-p", fmt.Sprintf("%d:%d", mindcli.config.ServeMPKPort, mindcli.config.ServeMPKPort),
		"-e", fmt.Sprintf("MIND_USER=%s", mindUser()),
		"-e", fmt.Sprintf("SERVE_REMOTE_PORT=%d", mindcli.config.ServeRemotePort),
		"-e", fmt.Sprintf("SERVE_MPK_PORT=%d", mindcli.config.ServeMPKPort),
		"-e", fmt.Sprintf("USER_HASH=%s", mindcli.userConfig.userHash),
		"-e", fmt.Sprintf("SKILLDIR=%s", mindcli.config.ContainerSkillDir),
		"-v", mindcli.skillMountPoint(),
		fmt.Sprintf("%s", mindcli.config.Image),
	}
	return append(c, args...)
}

func (mindcli *MindCli) execDocker(args []string) {
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func (mindcli *MindCli) execDockerOutput(args []string) string {
	var outb bytes.Buffer
	cmd := exec.Command("docker", args...)
	cmd.Stdout = &outb
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	return outb.String()
}

func (mindcli *MindCli) UpgradeImage() {
	mindcli.execDocker(mindcli.dockerUpgradeImageArgs())
}

func (mindcli *MindCli) X(args ...string) {
	mindcli.execDocker(mindcli.dockerXArgs(args...))
}

func (mindcli *MindCli) XOutput(args ...string) string {
	return mindcli.execDockerOutput(mindcli.dockerXArgs(args...))
}

func (mindcli *MindCli) RunSkill(noInstall bool, robotIp string, args ...string) {
	if robotIp == "" {
		var robot *Robot
		if len(args) < 1 {
			robot = mindcli.RobotByName(mindcli.DefaultRobotName())
		} else {
			robot = mindcli.RobotByName(args[0])
		}
		if robot == nil {
			fmt.Println("Could not find robot with that name\nPlease run `mind scan`")
			os.Exit(-1)
		}
		robots, _ := mindcli.RobotScanner.ScanIP(robot.IP, time.Duration(2))
		if len(robots) < 1 {
			fmt.Println("Could not find any robots\nPlease run `mind scan`")
			os.Exit(-1)
		}
		if robots[0].Name != robot.Name {
			fmt.Println("Looks like robot has changed IP\nPlease run `mind scan`")
			os.Exit(-1)
		}

		robotIp = robot.IP
	}
	ip, _ := GetLocalIPByNeighbourIP(robotIp)
	xParams := []string{"mindcli-run"}
	if noInstall {
		xParams = append(xParams, "-n")
	}
	if ret := net.ParseIP(robotIp); ret == nil {
		fmt.Println("The IP address you just typed does NOT meet the standard ipv4 format.")
		return
	}
	xParams = append(xParams,
		"skill.mpk",
		fmt.Sprintf("http://%s:%d", ip.String(), mindcli.config.ServeMPKPort),
		fmt.Sprintf("%v", robotIp),
	)
	mindcli.X(append(xParams, args...)...)

}

func (mindcli *MindCli) Login(email, pass string) error {
	userhash := strings.Trim(mindcli.XOutput(append([]string{"mindcli-login"}, email, pass)...), "\r\n")
	if len(userhash) == 0 {
		return errors.New("Login failed")
	}
	mindcli.userConfig.userHash = userhash
	return mindcli.userConfig.Write()
}

func (mindcli *MindCli) Scan(waitDuration int, args ...string) ([]Robot, error) {
	var robots []Robot
	var err error
	if len(args) < 1 {
		robots, err = mindcli.RobotScanner.BroadcastToNetwork(time.Duration(waitDuration))
		if err != nil {
			return nil, err
		}
	} else {
		robots, err = mindcli.RobotScanner.ScanIP(args[0], time.Duration(waitDuration))
		if err != nil {
			return nil, err
		}
	}
	mindcli.userConfig.Robots = robots
	return robots, mindcli.userConfig.Write()
}

func (mindcli *MindCli) RobotByName(robotName string) *Robot {
	for _, robot := range mindcli.userConfig.Robots {
		if robot.Name == robotName {
			return &robot
		}
	}
	return nil
}

func (mindcli *MindCli) SetDefaultRobotName(robotName string) error {
	for _, robot := range mindcli.userConfig.Robots {
		if robot.Name == robotName {
			mindcli.userConfig.DefaultRobotName = robotName
			mindcli.userConfig.Write()
			return nil
		}
	}
	return errors.New("Could not find robot with that name")
}

func (mindcli *MindCli) DefaultRobotName() string {
	return mindcli.userConfig.DefaultRobotName
}

func (mindcli *MindCli) DefaultRobotIP() string {
	robot := mindcli.RobotByName(mindcli.userConfig.DefaultRobotName)
	if robot == nil {
		return ""
	}
	return robot.IP
}

func (mindcli *MindCli) RunFlightTest(args ...string) {
	ips, err := GetLocalIPs()
	if err != nil || len(ips) < 1 {
		fmt.Println("Cannot find a local IP, check your network settings.")
		return
	}
	mindcli.X(append([]string{
		"mindcli-flighttest",
		"skill.mpk",
		fmt.Sprintf("%v", ips[0]),
		fmt.Sprintf("%v", mindcli.config.ServeMPKPort),
	}, args...)...)
}
