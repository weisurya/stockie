package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log logger

type logger struct {
	log *logrus.Logger
}

func NewLog() (l logger) {

	l.log = logrus.New()

	file, err := os.OpenFile(fmt.Sprintf("log_%s.log", time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.log.Out = os.Stdout
	} else {
		l.log.Out = file
	}

	return l
}

func (l logger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args)
}

func (l logger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args)
}
