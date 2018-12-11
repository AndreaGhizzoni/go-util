package network

import "net"

func GetIPFromInterfaceName(interfacee string) string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, i := range ifaces {
		if i.Name == interfacee {
			addrs, err := i.Addrs()
			if err != nil {
				return ""
			}
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					return v.IP.String()
				case *net.IPAddr:
					return v.IP.String()
				}
			}
		}
	}
	return ""
}
