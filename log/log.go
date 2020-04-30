package log

import (
  "fmt"
  "os"

  log "github.com/sirupsen/logrus"
)

func Init(level string) {
  //log.SetFormatter(&log.JSONFormatter{})
  log.SetFormatter(&log.TextFormatter{DisableColors: false, FullTimestamp: true, })
  log.SetOutput(os.Stdout)

  switch level {
  case "Debug":
    log.SetLevel(log.DebugLevel)
  case "Info":
    log.SetLevel(log.InfoLevel)
  case "Error":
    log.SetLevel(log.ErrorLevel)
  case "Fatal":
    log.SetLevel(log.FatalLevel)
  case "Panic":
    log.SetLevel(log.PanicLevel)
  }

}

func Debug(s string, f ...interface{}) {
  log.Debug(fmt.Sprintf(s, f...))
}

func Info(s string, f ...interface{}) {
  log.Info(fmt.Sprintf(s, f...))
}

func Error(s string, f ...interface{}) {
  log.Error(fmt.Sprintf(s, f...))
}

func Fatal(s string, f ...interface{}) {
  log.Fatal(fmt.Sprintf(s, f...))
  os.Exit(1)
}

func Panic(s string, f ...interface{}) {
  log.Panic(fmt.Sprintf(s, f...))
  os.Exit(1)
}
