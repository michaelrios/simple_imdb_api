package core

import (
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

type Dependencies struct {
	DB     SQL
	logger *logrus.Logger
}

func (deps *Dependencies) SetBaseLogger(logger *logrus.Logger) {
	deps.logger = logger
}

func (deps *Dependencies) GetLogger() *logrus.Entry {
	return deps.logger.WithField("requestID", ksuid.New().String())
}
