package nets

import (
	"net"
)

// GetIP 获取IP地址.
func GetIP() []byte {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err == nil {
		localAddr, ok := conn.LocalAddr().(*net.UDPAddr)

		if ok {
			return localAddr.IP
		}
	}

	return []byte{0, 0, 0, 0}
}
