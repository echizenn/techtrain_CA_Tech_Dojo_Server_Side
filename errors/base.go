package errors

import (
	"fmt"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

type BaseError struct {
	statusCode int
	level      log.Level
	msg        string
}

func (e BaseError) Error() string {
	return fmt.Sprintf("%s: code=%d, msg=%s", e.level, e.statusCode, e.msg)
}
