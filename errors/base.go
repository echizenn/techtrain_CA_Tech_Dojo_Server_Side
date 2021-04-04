package errors

import (
	"fmt"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

type BaseError struct {
	StatusCode int
	Level      log.Level
	Msg        string
}

func (e BaseError) Error() string {
	return fmt.Sprintf("%s: code=%d, M=%s", e.Level, e.StatusCode, e.Msg)
}
