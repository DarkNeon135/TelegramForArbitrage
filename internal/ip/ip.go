package ip

import (
	"fmt"
	"net"
)

func GetLocalAddress() (net.IP, error) {
	con, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, fmt.Errorf("error in getting a local IP address. Error: %s", err)
	}

	localAddress := con.LocalAddr().(*net.UDPAddr)

	return localAddress.IP, nil
}
