package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

// keyType is used to retrieve the logger from a value stored by a context.
type keyType struct{}

var Logger = logrus.StandardLogger()

func G(ctx context.Context) *logrus.Entry {
	if logger := ctx.Value(keyType{}); logger != nil {
		return logger.(*logrus.Entry)
	}
	return logrus.NewEntry(Logger)
}

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, keyType{}, logger)
}
