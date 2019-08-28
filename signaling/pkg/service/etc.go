package service

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)


var ServerIpAddress string

func CreateDirectory(filePath string) bool {
	dirName := path.Dir(filePath)
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		return false
	}

	return false
}
func ExternalIP() (string, error) {
	if len(ServerIpAddress) > 0 {
		return ServerIpAddress, nil
	}
	return ExternalIPRefresh()
}
func ExternalIPRefresh() (string, error) {
	TimeTrack(time.Now(), GetFunctionName())
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var ip net.IP
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
		}
	}
	if ip != nil {
		ServerIpAddress = ip.String()
		return ip.String(), nil
	} else {
		return "", errors.New("confirm you have no connection to the network")
	}
}
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Fprintf(os.Stderr, "%s took %v\n", name, elapsed.Nanoseconds())
}
func GetFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	tokens := strings.Split(f.Name(), ".")

	return tokens[len(tokens)-1]
}