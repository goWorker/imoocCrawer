package common

import (
	"net"
	"xlog"
)

var (
	localIP string
)

func GetLocalIP() (ip string,err error) {
	if len(localIP) > 0 {
		ip = localIP
		return
	}
	addrs,err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		ipAddr,ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipAddr.IP.IsLoopback() {
			//continue
			xlog.LogDebug("ip:%#v\n",ipAddr.IP.String())
			localIP = ipAddr.IP.String()
			ip = localIP
			return

		}

		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		//xlog.LogDebug("ip:%#v\n",ipAddr.IP.String())
		//localIP = ipAddr.IP.String()
		//ip = localIP
		//return

	}
	return

}
