package sys

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

type Info struct {
	OS              string              `json:"operating_system"`
	Hostname        string              `json:"hosting"`
	NetworkAdapters map[string][]string `json:"network"`
	EnvVars         map[string]string   `json:"environment_variables"`
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

func getEnvVars() map[string]string {
	out := make(map[string]string)
	for _, e := range os.Environ() {
		v := strings.Split(e, "=")
		out[v[0]] = v[1]
	}
	return out
}

func New() *Info {
	// determine hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return &Info{
		OS:              fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH),
		EnvVars:         getEnvVars(),
		Hostname:        hostname,
		NetworkAdapters: getIPAddresses(),
	}
}
