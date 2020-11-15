package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func Init(level string) {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: false, FullTimestamp: true})
	logrus.SetOutput(os.Stdout)

	switch level {
	case "Trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "Debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "Info":
		logrus.SetLevel(logrus.InfoLevel)
	case "Warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "Error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "Fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "Panic":
		logrus.SetLevel(logrus.PanicLevel)
	}

}

func Debug(s string, f ...interface{}) {
	logrus.Debug(fmt.Sprintf(s, f...))
}

func Info(s string, f ...interface{}) {
	logrus.Info(fmt.Sprintf(s, f...))
}

func Error(s string, f ...interface{}) {
	logrus.Error(fmt.Sprintf(s, f...))
}

func Fatal(s string, f ...interface{}) {
	logrus.Fatal(fmt.Sprintf(s, f...))
	os.Exit(1)
}

func Panic(s string, f ...interface{}) {
	s_ := fmt.Sprintf(s, f...)
	logrus.Panic(s_)
	panic(s_)
}
