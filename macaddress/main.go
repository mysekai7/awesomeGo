package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/appleboy/easyssh-proxy"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}

	fmt.Println("-----------")

	fmt.Println(getMacAddr2())
	fmt.Println("-----------")

	sshClient := easyssh.MakeConfig{
		User:     "root",
		Server:   "150.109.39.180",
		Port:     "8022",
		Password: "2CV1eVed9gzy",
		Timeout:  5 * time.Second,
	}
	_, err = sshClient.Connect()
	if err != nil {
		log.Fatalln("Connect:", err)
	}

	cmd := SSHCommand{sshClient}
	ok, err := cmd.IsRoot()
	if err != nil {
		log.Fatalln("isroot err: ", err)
	}
	if ok {
		fmt.Println("root is ok")
	}

	mac, err := cmd.GetMacAddress()
	if err != nil {
		log.Fatalln("get mac err: ", err)
	}
	fmt.Println("mac by address: ", mac)

	cpu, err := cmd.GetCPUID()
	if err != nil {
		log.Fatalln("get cpu err: ", err)
	}
	fmt.Println("cpu: ", cpu)
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
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

const (
	IsRootCmd                = "is_root"
	GetCPUIDCmd              = "get_cpu_id"
	GetSystemSerialNumberCmd = "get_system_serial_number"
	GetMacAddressCmd         = "get_mac_address"
)

var Command = map[string]string{
	GetCPUIDCmd:              `dmidecode -t 4 | grep ID | head -1`,
	GetSystemSerialNumberCmd: `dmidecode -s system-serial-number`,
	GetMacAddressCmd:         `cat /sys/class/net/$(route | grep default | awk '{print $NF}')/address`,
	IsRootCmd:                `whoami`,
}

type SSHCommand struct {
	easyssh.MakeConfig
}

func (cmd *SSHCommand) ShellCommand(shellCmd string) (string, error) {
	if executeCmd, ok := Command[shellCmd]; ok == false {
		return "", errors.New(fmt.Sprintf("cmd %s not exists", shellCmd))
	} else {
		stdout, stderr, isTimeout, err := cmd.Run(executeCmd, 5)
		if err != nil {
			if isTimeout == true {
				return "", errors.New(fmt.Sprintf("cmd %s is timeout, error is: %s", IsRootCmd, err.Error()))
			} else {
				return "", errors.New(fmt.Sprintf("cmd $s error, stderr: %s", IsRootCmd, stderr))
			}
		}
		return strings.TrimSpace(stdout), nil
	}
}

func (cmd *SSHCommand) IsRoot() (bool, error) {
	result, err := cmd.ShellCommand(IsRootCmd)
	if err != nil {
		return false, err
	}
	if result != "root" {
		return false, errors.New(fmt.Sprintf("please use administrator privileges"))
	}
	return true, nil
}

func (cmd *SSHCommand) GetCPUID() (string, error) {
	result, err := cmd.ShellCommand(GetCPUIDCmd)
	if err != nil {
		return "", err
	}
	slice := strings.Split(result, ":")
	if len(slice) != 2 {
		return "", nil
	}
	return strings.TrimSpace(slice[1]), nil
}

func (cmd *SSHCommand) GetSystemSerialNumber() (string, error) {
	result, err := cmd.ShellCommand(GetSystemSerialNumberCmd)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (cmd *SSHCommand) GetMacAddress() (string, error) {
	result, err := cmd.ShellCommand(GetMacAddressCmd)
	if err != nil {
		return "", err
	}
	return result, nil
}
