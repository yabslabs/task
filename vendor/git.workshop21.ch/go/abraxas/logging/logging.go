package logging

import (
	"git.workshop21.ch/go/abraxas/utils/pairs"
	"github.com/sirupsen/logrus"
)

const (
	fieldLogID = "abraxasLogID"
)

func WithError(id string, err error) *logrus.Entry {
	return WithID(id).WithError(err)
}

func WithID(id string) *logrus.Entry {
	return logrus.WithField(fieldLogID, id)
}

func WithIDFields(id string, fields ...interface{}) *logrus.Entry {
	m := pairs.Pairs(fields...)
	m[fieldLogID] = id
	return logrus.WithFields(m)
}
