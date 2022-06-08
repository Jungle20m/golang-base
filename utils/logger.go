package utils

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLoggerInstance(format, output, level int, directory, fileName string) (*logrus.Logger, error) {
	log := logrus.New()

	// Cài đặt định dạng logger
	switch format {
	case 1:
		formatter := logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		log.SetFormatter(&formatter)
	default:
		formatter := new(logrus.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.FullTimestamp = true
		log.SetFormatter(formatter)
	}

	// Cài đặt output logger
	switch output {
	case 1:
		log.SetOutput(&lumberjack.Logger{
			Filename:   filepath.Join(directory, fileName),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,    //days
			Compress:   false, // disabled by default
		})
	default:
		// terminal
	}

	// Cài đặt level hiển thị
	var lv = logrus.PanicLevel
	switch level {
	case 1:
		lv = logrus.FatalLevel
	case 2:
		lv = logrus.ErrorLevel
	case 3:
		lv = logrus.WarnLevel
	case 4:
		lv = logrus.InfoLevel
	case 5:
		lv = logrus.DebugLevel
	case 6:
		lv = logrus.TraceLevel
	default:
		lv = logrus.PanicLevel
	}
	log.SetLevel(lv)

	return log, nil
}
