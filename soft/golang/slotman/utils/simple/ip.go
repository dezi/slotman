package simple

import (
	"errors"
	"net"
	"sort"
	"strings"
)

func GetLocalIPV4() (ip net.IP, err error) {
	ip, err = GetLocalIP(false, nil)
	return
}

func GetLocalIP(ipv6 bool, suffixes []string) (ip net.IP, err error) {

	iFaces, err := net.Interfaces()
	if err != nil {
		return
	}

	//
	// Sort interfaces to always get the
	// same ip if multiple interface are
	// present. Also prefer en(x) over wl(x).
	//

	sort.Slice(iFaces, func(i, j int) bool {
		return iFaces[i].Name < iFaces[j].Name
	})

	for _, iface := range iFaces {

		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		var addrs []net.Addr
		addrs, err = iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {

			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipv6 {

				if ipNet.IP.To4() == nil {

					if suffixes == nil {
						ip = ipNet.IP
						return
					}

					for _, suffix := range suffixes {
						if strings.HasSuffix(ipNet.IP.String(), suffix) {
							ip = ipNet.IP
							return
						}
					}
				}

			} else {
				if ipNet.IP.To4() != nil {
					ip = ipNet.IP
					return
				}
			}
		}
	}

	err = errors.New("no local ip")
	return
}
