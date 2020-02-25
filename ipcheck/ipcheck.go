package ipcheck

import (
	"device-monitor-api/constants"
	"fmt"
	"net"
)

func GetIPCountry(ipAddress string) (country string) {

	ipA := net.ParseIP(ipAddress)

	for _, c := range constants.AmericanCIDR {
		_, ipnet, err := net.ParseCIDR(c)
		if err != nil {
			fmt.Println("GetIPCountry : America, %s", err.Error())
		}
		if ipnet.Contains(ipA) {
			country = "America"
			return
		}
	}

	for _, c := range constants.JapanCIDR {
		_, ipnet, err := net.ParseCIDR(c)
		if err != nil {
			fmt.Println("GetIPCountry : Japan, %s", err.Error())
		}
		if ipnet.Contains(ipA) {
			country = "Japan"
			return
		}
	}
	country = "Unknown"
	return
}
