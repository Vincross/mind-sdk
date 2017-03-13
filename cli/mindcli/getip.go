package mindcli

import (
	"errors"
	"net"
)

func GetLocalIPByNeighbourIP(neighbourIPStr string) (net.IP, error) {
	neighbourIP := net.ParseIP(neighbourIPStr)
	if neighbourIP == nil {
		return nil, errors.New("Could not parse provided IP")
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok &&
			!ip.IP.IsLoopback() &&
			ip.IP.To4() != nil &&
			ip.Contains(neighbourIP) {
			return ip.IP, nil
		}
	}
	return nil, errors.New("Could not find local IP")
}

func GetLocalIPs() ([]net.IP, error) {
	ips := []net.IP{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok &&
			!ip.IP.IsLoopback() &&
			ip.IP.To4() != nil {
			ips = append(ips, ip.IP)
		}
	}
	return ips, nil
}
