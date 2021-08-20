package logger

import (
	"fmt"
	"io"
	logging "log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {

	log = logrus.New()
	logrus.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{CallerPrettyfier: caller(),
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "caller",
		}})
	log.Level = logrus.DebugLevel
	date := time.Now().Format("2006-01-02")
	createFolder()
	fileName := fmt.Sprintf("logs/%v.log", date)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	log.SetReportCaller(false)

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

}

func createFolder() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			logging.Fatalf("Error creating logs folder %v", err)
		}
	}
}

func caller() func(*runtime.Frame) (function string, file string) {
	return func(f *runtime.Frame) (function string, file string) {
		p, _ := os.Getwd()

		return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
	}
}

func Info(msg string) {
	log.Info(msg)
}
func Warn(msg string) {
	log.Warn(msg)
}
func Error(err error) {
	log.Error(err)
}
func Debug(arg ...interface{}) {
	log.Debug(arg)
}
func Fatal(err error) {
	log.Fatal(err)
}
