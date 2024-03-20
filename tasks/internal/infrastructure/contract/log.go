package contract

import (
	"github.com/sirupsen/logrus"
)

type Log interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	WithField(string, interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
}
