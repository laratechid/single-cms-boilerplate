package helper

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogInfo(msg string, stack ...string) {
	var msgDetail string = msg
	if stack != nil {
		if stack[0] != "" {
			msgDetail = fmt.Sprintf("src: %s | info: %s", stack[0], msg)
		}
	}
	logrus.Info(msgDetail)
}

func LogErr(err error, stacks ...string) {
	if stacks != nil {
		if stacks[0] != "" {
			stack := stacks[0]
			logrus.WithFields(logrus.Fields{
				"stack": stack,
			}).Error(err.Error())
			return
		}
	}
	logrus.Error(err.Error())
}

func LogErrString(err string, stack ...string) {
	msgDetail := fmt.Sprintf("src: %s | err: %s", stack[0], err)
	logrus.Error(msgDetail)
}
