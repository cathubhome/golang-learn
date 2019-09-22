package main

import (
	"fmt"
	"net"
)

var (
	localIpArray []string
)

/**
获取本地IP
*/
func init() {
	//获取所有网卡
	addrs, e := net.InterfaceAddrs()
	if e != nil {
		panic(fmt.Sprintf("get local ip failed,%v\n", e))
	}
	for _, addr := range addrs {
		//取网络地址的网卡的信息并检查IP地址是否是回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			//能正常转成ipv4
			if ipnet.IP.To4() != nil {
				localIpArray = append(localIpArray, ipnet.IP.String())
			}
		}
	}

	fmt.Println("localIpArray", localIpArray)
}
