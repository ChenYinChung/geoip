package geolib

import (
	"archive/tar"
	"compress/gzip"
	"geoip_auto/logger"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// GeoIp mmdb所在位置，必需置換license_key
const url = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=licenseKey&suffix=tar.gz"

// 檔案名稱
const geoLite2CityFilename = "GeoLite2-City.tar.gz"

var log = logger.NewLogger()

// GeoIp License Key
type GeoIpLicenseKey struct {
	key string
}

// GeoIP主體
type GeoIP struct {
	LicenseKey GeoIpLicenseKey
	extract    string // 解壓縮路徑
	mux        sync.RWMutex
}

func New(extrapath string, licenseKey GeoIpLicenseKey) *GeoIP {
	return &GeoIP{extract: extrapath, LicenseKey: licenseKey}
}

func (geoip *GeoIP) RemoveAndCreateGeoLite2City() {
	os.RemoveAll(geoip.extract)
	if err := os.Mkdir(geoip.extract, 0755); err != nil {
		log.Fatalf("data: Mkdir() failed: %s", err.Error())
	}
}

// DownloadFile will download from a given url to a file. It will
// write as it downloads (useful for large files).
func (geoip *GeoIP) DownloadFileGeoLite2City() error {

	// Get the data
	resp, err := http.Get(strings.Replace(url, "licenseKey", geoip.LicenseKey.key, 1))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(geoip.extract + geoLite2CityFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (geoip *GeoIP) ExtractGeoLite2CityTarGz() {

	gzipStream, err := os.Open(geoip.extract + geoLite2CityFilename)
	if err != nil {
		log.Println("error")
	}

	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			log.Println("TypeDir header.Name:", header.Name)
			//if err := os.Mkdir(geoip.extract+header.Name, 0755); err != nil {
			//	logger.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			//}
		case tar.TypeReg:

			//fmt.Println("TypeReg header.Name:", header.Name)

			if strings.HasSuffix(header.Name, "mmdb") {
				strs := strings.Split(header.Name, "/")
				log.Println("TypeReg header.Name:", header.Name, " strs:", strs)
				fn := strs[1]

				outFile, err := os.Create(geoip.extract + fn)

				if err != nil {
					log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
				}
				if _, err := io.Copy(outFile, tarReader); err != nil {
					log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
				}
				outFile.Close()
			}

		default:
			log.Fatalf("ExtractTarGz: uknown type: %b in %s", header.Typeflag, header.Name)
		}
	}
}

type Logger struct {
	LicenseKey string `yaml:"licenseKey"`
	ExtraPath  string `yaml:"extraPath"`
}

func NewGeoConfig() *Logger {

	vip := viper.New()
	vip.SetConfigName("geo")
	vip.SetConfigType("yml")
	vip.AddConfigPath("../config")
	err := vip.ReadInConfig()
	if err != nil {
		log.Panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}

	logConf := &Logger{}
	err = vip.Unmarshal(&logConf)

	if err != nil {
		log.Panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}

	//fmt.Println(logConf)

	return logConf
}

func init() {
	var log = logger.NewLogger()
	log.Info("Initial Timer Here")
}
