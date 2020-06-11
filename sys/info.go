package sys

import (
	"net"
	"os"
)

type Info struct {
	OS          string
	Hostname    string
	IPAddresses map[string][]string
}

func getIPAddresses() (out map[string][]string) {
	ifaces, err := net.Interfaces()
	if err != nil {

	}

	out = make(map[string][]string)

	for _, i := range ifaces {
		var ips []string

		addrs, err := i.Addrs()

		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip.String())
		}

		out[i.Name] = ips
	}

	return
}

func New() *Info {
	// determine hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// determine ip addresses

	return &Info{
		OS:          "mac",
		Hostname:    hostname,
		IPAddresses: getIPAddresses(),
	}
}
