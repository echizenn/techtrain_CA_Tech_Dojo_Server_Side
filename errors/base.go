package errors

import (
	"fmt"
	"time"

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
			logger.Debug("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case Info:
			logger, _ := zap.NewDevelopment()
			logger.Info("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case Warn:
			logger, _ := zap.NewDevelopment()
			logger.Warn("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case Error:
			logger, _ := zap.NewDevelopment()
			logger.Error("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case DPanic:
			logger, _ := zap.NewDevelopment()
			logger.DPanic("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case Panic:
			logger, _ := zap.NewDevelopment()
			logger.Panic("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		case Fatal:
			logger, _ := zap.NewDevelopment()
			logger.Fatal("zap", zap.String("msg", err.Error()), zap.Time("now", time.Now()))
		}
	}
}
