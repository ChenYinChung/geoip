# Geoip 

## 目录结构
* conf：配置文件
* data：Geoip mmdb
  * GeoLite2-City.mmdb
  * GeoLite2-City.tar.gz
* geolib： geo共用
* iplib: ip期用
* logger： 日誌
* verify：驗證ip & geoip

## 下載GeoLite2-City.mmdb

```
conf := NewGeoConfig()
renew := New(conf.ExtraPath, GeoIpLicenseKey{key: conf.LicenseKey})

renew.RemoveAndCreateGeoLite2City()
e := renew.DownloadFileGeoLite2City()
if e != nil {
	log.Fatal(e)
}
renew.ExtractGeoLite2CityTarGz()
```
