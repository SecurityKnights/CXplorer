package modules

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PortScan(ipAddress string, protocol string) bool {
	conn, err := net.DialTimeout(protocol, ipAddress, 3*time.Second)

	if err != nil {
		return false
	} else {
		defer conn.Close()
		return true
	}
}

func PortScanner(ipAddress string, protocol string, startPort, endPort int) {
	defer wg.Done()

	for i := startPort; i <= endPort; i++ {
		temp := net.JoinHostPort(ipAddress, strconv.Itoa(i))

		if PortScan(temp, protocol) {
			fmt.Printf("Port %d is open on %s\n", i, ipAddress)
		}
	}
}

func InitPortScanner(ipAddress string, protocol string, portRange string) {
	ports := strings.Split(portRange, "-")

	for i := range ports {
		ports[i] = strings.TrimSpace(ports[i])
	}

	if len(ports) > 2 {
		fmt.Println("Wrong Input! Kindly refer the help menu using -h flag")
		os.Exit(1)
	}

	startPort, _ := strconv.Atoi(ports[0])
	endPort, _ := strconv.Atoi(ports[1])

	fmt.Println("Target:\t\t\t", ipAddress)
	fmt.Println("OS:\t\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Protocol:\t\t", protocol)

	fmt.Println("=============================================================")

	fmt.Println("Scanning Started at", time.Now().Format("2006-01-02 3:4:5"))

	fmt.Println("=============================================================")

	wg.Add(1)

	timeStart := time.Now()

	go PortScanner(ipAddress, protocol, startPort, endPort)

	wg.Wait()

	timeElapsed := time.Since(timeStart)

	fmt.Println("=============================================================")

	fmt.Println("Scanning Done")
	fmt.Printf("Time Taken: %s\n", timeElapsed)
}
