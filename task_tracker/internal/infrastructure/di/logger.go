package di

import (
	"os"

	"github.com/sirupsen/logrus"

	"async-arch/task_tracker/internal/infrastructure/contract"
)

func NewLogger() contract.Log {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
