package verify

import (
	"geoip_auto/geolib"
	"geoip_auto/logger"
	"github.com/oschwald/geoip2-golang"
	"net"
	"sync"
)

const filename = "GeoLite2-City.mmdb"

var once sync.Once
var instance *GeoVerify

type GeoVerify struct {
}

//type City struct {
//	City []struct {
//		Country string
//		Name    string
//	}
//}

func GetInstance() *GeoVerify {
	once.Do(func() {
		instance = &GeoVerify{}
	})
	return instance
}

func (geo *GeoVerify) Parse(ip net.IP) {
	conf := geolib.NewGeoConfig()
	log := logger.NewLogger()
	db, err := geoip2.Open(conf.ExtraPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil

	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	if len(record.Subdivisions) > 0 {
		log.Printf("English subdivision name: %v", record.Subdivisions[0].Names["en"])
		log.Printf("English subdivision name: %v", record.Subdivisions[0].Names["zh-CN"])
	}

	log.Infof("ISO country code: %v", record.Country.IsoCode)
	//fmt.Printf("Taiwan country name: %v\n", record.Country.Names["tw"])
	//log.Infof("Time zone: %v", record.Location.TimeZone)
	//log.Infof("Coordinates: %v, %v", record.Location.Latitude, record.Location.Longitude)

	log.Infof("國家[zh-CN]:[%s]", record.Country.Names["zh-CN"])
	log.Infof("國家[en]:[%s]", record.Country.Names["en"])

	log.Infof("record.City.Names[zh-CN]:[%s]", record.City.Names["zh-CN"])
	log.Infof("record.City.Names[en]:[%s]", record.City.Names["en"])
	//for k, _ := range record.City.Names {
	//
	//	log.Infof("record.City.Names[%s]:[%s]", k, record.City.Names[k])
	//	record.Country.IsoCode
	//
	//}

	//log.Info(record)

}
