package main

import (
	"bytes"
	"fmt"
	"net"

	net2 "github.com/shirou/gopsutil/net"
)

func main() {

	fmt.Println("-----------")

	fmt.Println(getMacAddr2())
	fmt.Println("-----------")

	net2.Addr{}
	net2.InterfaceAddr{}
	net2.

}

// getMacAddr gets the MAC hardware
// address of the host machine
func getMacAddr2() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}
