package vlog

import (
	"github.com/sirupsen/logrus"
	"os"
)

var ErrorLog *logrus.Logger
var AccessLog *logrus.Logger
var errorLogFile = "./tmp/log/error.log"
var accessLogFile = "./tmp/log/access.log"
func init () {
	initErrorLog()
	initAccessLog()
}

func initErrorLog() {
	ErrorLog = logrus.New()
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})
	file , err := os.OpenFile(errorLogFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	ErrorLog.SetOutput(file)
}

func initAccessLog() {
	AccessLog = logrus.New()
	AccessLog.SetFormatter(&logrus.JSONFormatter{})
	file , err := os.OpenFile(accessLogFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	AccessLog.SetOutput(file)
}