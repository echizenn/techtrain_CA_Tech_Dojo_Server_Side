package errors

import (
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type BaseError struct {
	StatusCode int
	Level      Level
	Msg        string
}

func (e BaseError) Error() string {
	return fmt.Sprintf("%s: code=%d, M=%s", e.Level, e.StatusCode, e.Msg)
}

func EmitLog(err error) {
	var baseError *BaseError
	if xerrors.As(err, &baseError) {
		switch baseError.Level {
		case Debug:
			logger, _ := zap.NewDevelopment()
			logger.Info("Hello zap", zap.String("key", "value"), zap.String("now", "2021-04-04"))
		case Info:
			// Info
		case Warn:
			// Warn
		case Error:
			// Error
		case DPanic:
			// DPanic
		case Panic:
			// Panic
		case Fatal:
			// Fatal
		}
	}
}
