package errors

import (
	"fmt"
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
	return
}
