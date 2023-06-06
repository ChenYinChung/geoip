package geolib

import (
	"testing"
)

func TestGetGeoDatabase(t *testing.T) {
	conf := NewGeoConfig()
	renew := New(conf.ExtraPath, GeoIpLicenseKey{key: conf.LicenseKey})

	renew.RemoveAndCreateGeoLite2City()
	e := renew.DownloadFileGeoLite2City()
	if e != nil {
		log.Fatal(e)
	}

	renew.ExtractGeoLite2CityTarGz()
}
