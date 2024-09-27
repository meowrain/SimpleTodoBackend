package ips

import (
	"net"
	"todoBackend/utils/loggers"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	loggers.TodoLogger.Infoln("addrs: ", addrs)
	if err != nil {
		loggers.TodoLogger.Errorf("无法获取本机IP地址:%s", err)
		return ""
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			resIp := ipNet.IP.To4()
			if resIp != nil {
				return resIp.String()
			}
		}
	}
	return ""
}
