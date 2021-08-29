package logger

import (
	"fmt"
	"io"
	logging "log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/go4digital/booknow-api/constants"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	// First create logs folder if not exist
	createFolder()

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{CallerPrettyfier: caller(),
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "caller",
		}})
	date := time.Now().Format(constants.YYYYMMDD)
	fileName := fmt.Sprintf("%v/%v.log", constants.LOGS_FOLDER_NAME, date)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, constants.READ_WRITE)

	if err != nil {
		logging.Fatalf("Error opening file:- %v", err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
}

func createFolder() {
	if _, err := os.Stat(constants.LOGS_FOLDER_NAME); os.IsNotExist(err) {
		err = os.Mkdir(constants.LOGS_FOLDER_NAME, constants.ADMIN_RIGHT)
		if err != nil {
			logging.Fatalf("Error creating logs folder:- %v", err)
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
