package nets

import (
	"net"
)

// GetIP 获取IP地址.
func GetIP() []byte {
	if conn, err := net.Dial("udp", "8.8.8.8:53"); err == nil {
		if localAddr, ok := conn.LocalAddr().(*net.UDPAddr); ok {
			return localAddr.IP
		}
	}

	return []byte{0, 0, 0, 0}
}
