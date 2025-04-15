package helper

import (
	"fmt"

	"github.com/go-stack/stack"
	"github.com/sirupsen/logrus"
)

func LogInfo(msg string) {
	fn := stack.Caller(1).Frame().Function
	msgDetail := fmt.Sprintf("%s | %s", fn, msg)
	logrus.Info(msgDetail)
}

func LogErr(err error) {
	fn := stack.Caller(1).Frame().Function
	msgDetail := fmt.Sprintf("%s | %s", fn, err.Error())
	logrus.Error(msgDetail)
}

func LogErrString(err string) {
	fn := stack.Caller(1).Frame().Function
	msgDetail := fmt.Sprintf("%s | %s", fn, err)
	logrus.Error(msgDetail)
}
