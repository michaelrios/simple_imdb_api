package core

import (
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	DB     DB
	Logger *logrus.Logger
}

func (deps *Dependencies) SetBaseLogger(logger *logrus.Logger) {
	deps.Logger = logger
}

func (deps *Dependencies) GetRequestLogger() *logrus.Entry {
	return deps.Logger.WithField("requestID", ksuid.New().String())
}
