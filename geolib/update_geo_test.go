package geolib

import (
	"testing"
)

func TestGetGeoDatabase(t *testing.T) {
	conf := NewGeoConfig()
	updater := New(conf.ExtraPath, GeoIpLicenseKey{key: conf.LicenseKey})

	updater.RemoveAndCreateGeoLite2City()
	e := updater.DownloadFileGeoLite2City()
	if e != nil {
		log.Fatal(e)
	}

	updater.ExtractGeoLite2CityTarGz()
}
