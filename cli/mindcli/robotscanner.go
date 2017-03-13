package mindcli

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Robot struct {
	Name string `json:"Name"`
	IP   string `json:"IP"`
}

type RobotScannerConfig struct {
	Message string
	Port    int
}

type RobotScanner struct {
	robots []Robot
	Config *RobotScannerConfig
}

func NewRobotScanner(config *RobotScannerConfig) *RobotScanner {
	return &RobotScanner{
		Config: config,
	}
}

func (scanner *RobotScanner) HasRobot(robot Robot) bool {
	for _, r := range scanner.robots {
		if r == robot {
			return true
		}
	}
	return false
}

func IsTCPPortAvailable(port int) bool {
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func IsUDPPortAvailable(port int) bool {
	conn, err := net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	conn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func (scanner *RobotScanner) server() {
	// For some reason ListenUDP is not returning error
	// when port is already in use so we have to check manually
	if !IsUDPPortAvailable(scanner.Config.Port) || !IsTCPPortAvailable(scanner.Config.Port) {
		fmt.Printf("Port %d already in use\n", scanner.Config.Port)
		os.Exit(-1)
	}
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: scanner.Config.Port,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()
	for {
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg := strings.TrimSpace(string(data[:read]))
		if msg == scanner.Config.Message {
			continue
		}
		if !strings.Contains(msg, "|") {
			continue
		}
		name := strings.SplitN(msg, "|", 2)[1]
		robot := Robot{Name: name, IP: remoteAddr.IP.String()}
		if !scanner.HasRobot(robot) {
			scanner.robots = append(scanner.robots, robot)
		}
	}
}

func (scanner *RobotScanner) sendMsg(addr string) error {
	serverAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	localAddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	buf := []byte(scanner.Config.Message)
	conn.Write(buf)
	return nil
}

func (scanner *RobotScanner) ScanNetwork(waitDuration time.Duration) ([]Robot, error) {
	scanner.robots = []Robot{}
	ips, err := GetLocalIPs()
	if err != nil {
		return nil, err
	}
	go scanner.server()
	for _, ip := range ips {
		ipparts := strings.Split(ip.String(), ".")
		for i := 1; i < 255; i += 1 {
			sendToIP := fmt.Sprintf("%s.%s.%s.%d:%d", ipparts[0], ipparts[1], ipparts[2], i, scanner.Config.Port)
			_ = scanner.sendMsg(sendToIP)
		}
	}
	time.Sleep(waitDuration * time.Second)
	return scanner.robots, nil
}

func (scanner *RobotScanner) BroadcastToNetwork(waitDuration time.Duration) ([]Robot, error) {
	scanner.robots = []Robot{}
	ips, err := GetLocalIPs()
	if err != nil {
		return nil, err
	}
	go scanner.server()
	for _, ip := range ips {
		ipparts := strings.Split(ip.String(), ".")
		for i := 1; i < 20; i += 1 {
			sendToIP := fmt.Sprintf("%s.%s.%s.%d:%d", ipparts[0], ipparts[1], ipparts[2], 255, scanner.Config.Port)
			_ = scanner.sendMsg(sendToIP)
		}
	}
	time.Sleep(waitDuration * time.Second)
	return scanner.robots, nil
}

func (scanner *RobotScanner) ScanIP(ip string, waitDuration time.Duration) ([]Robot, error) {
	scanner.robots = []Robot{}
	go scanner.server()
	sendToIP := fmt.Sprintf("%s:%d", ip, scanner.Config.Port)
	_ = scanner.sendMsg(sendToIP)
	time.Sleep(waitDuration * time.Second)
	return scanner.robots, nil
}
