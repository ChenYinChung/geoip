package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewLogger() *logrus.Logger {
	var log = logrus.New()
	//log輸出為json格式
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	//輸出設定為標準輸出(預設為stderr)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	l := NewLoggerConfig()
	newLevel, err := logrus.ParseLevel(l.Level)
	if err != nil {
		//設定要輸出的log等級
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(newLevel)
	}

	return log
}

type Logger struct {
	Level string `yaml:"level"`
}

func NewLoggerConfig() *Logger {

	vip := viper.New()
	vip.SetConfigName("logger")
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
