package utils

import (
	"os"
	"path"

	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	logPath = "./logs/"
	logFile = "api.log"
)
var LogInstance = logrus.New()

// init log
func init() {
	logFileName := path.Join(logPath, logFile)

	rolling(logFileName)

	LogInstance.SetFormatter(&logrus.TextFormatter{})

	loglevel := os.Getenv("LOG_LEVEL")

	LogInstance.SetLevel(logrus.DebugLevel) // default log level: DEBUG
	if len(loglevel) != 0 {
		loglevel = strings.ToUpper(strings.Trim(loglevel, " "))
		switch {
		case loglevel == "INFO":
			LogInstance.SetLevel(logrus.InfoLevel)
		case loglevel == "WARN":
			LogInstance.SetLevel(logrus.WarnLevel)
		case loglevel == "ERROR":
			LogInstance.SetLevel(logrus.ErrorLevel)
		case loglevel == "FATAL":
			LogInstance.SetLevel(logrus.FatalLevel)
		}
	}
}

func rolling(logFile string) {
	// output setup
	LogInstance.SetOutput(&lumberjack.Logger{
		Filename:   logFile, //log file path
		MaxSize:    512,     // Sinle file size
		MaxBackups: 10,      // total number of log files
		MaxAge:     7,       // rotation period, unit: day
		Compress:   true,    // need to compress log file?

	})

}
