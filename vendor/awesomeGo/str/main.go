package main

import (
	"strings"
	"fmt"
)

func main() {

	hosts := "192.168.0.1,192.168.0.2,192.168.0.3"
	ports := "8080,8081,8082"


	res := getHosts(hosts, ports)

	fmt.Println("res: %v", res)
	
}


func getHosts(host, port string) []string {
	var add []string
	host = strings.TrimSpace(host)
	if len(host) > 0 {
		if strings.Contains(host, ",") {
			ports := strings.Split(port, ",")
			hosts := strings.Split(host, ",")
			for i, h := range hosts{
				add = append(add, h + ":" + ports[i])
			}
		} else {
			add = []string{host + ":" + port}
		}
	}
	return add
}