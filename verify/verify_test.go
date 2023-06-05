package verify

import (
	"geoip_auto/logger"
	"net"
	"testing"
)

func TestOpenGeoDatabase(t *testing.T) {
	v := GetInstance()
	//60.251.54.196,114.44.128.242

	v.Parse(net.ParseIP("60.251.54.196"))
	v.Parse(net.ParseIP("114.44.128.242"))
	v.Parse(net.ParseIP("23.255.19.64"))

}

func TestCIDR(t *testing.T) {
	log := logger.NewLogger()
	network := "192.168.5.0/24"
	clientips := []string{
		"192.168.5.1",
		"192.168.6.0",
	}
	_, subnet, _ := net.ParseCIDR(network)
	for _, clientip := range clientips {
		ip := net.ParseIP(clientip)
		if subnet.Contains(ip) {
			log.Println("IP in subnet", clientip)
		} else {
			log.Println("IP not in subnet", clientip)
		}
	}
}
