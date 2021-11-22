package helpers

import (
	"errors"
	"net"
)

// FindIP - used for get external IP address
func FindIP(filterFunc FilterIP) (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, netInterface := range netInterfaces {
		// When interface is down
		if netInterface.Flags&net.FlagUp == 0 {
			continue
		}

		// When it is loopBack interface
		if netInterface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := netInterface.Addrs()
		if err != nil {
			return "", err
		}

		address, err := filterFunc(addresses)
		if err == nil {
			return address, nil
		}
	}

	return "", errors.New("problems with address detection")
}

// FilterIP - used as contract for all filters functions
type FilterIP func(addresses []net.Addr) (string, error)

// FilterIPv4 - specific implementation of address filtering
func FilterIPv4(addresses []net.Addr) (string, error) {
	for _, addr := range addresses {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}

		// When it's loopBack address
		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			// When it's not IPv4 address
			continue
		}

		return ip.String(), nil
	}

	return "", errors.New("error with find IPv4 address")
}

// FilterIPv4Stub - used as stub for testing
func FilterIPv4Stub(addresses []net.Addr) (string, error) {
	return "192.168.1.2", nil
}
