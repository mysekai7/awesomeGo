package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func main() {

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hostStat.HostID)

}
