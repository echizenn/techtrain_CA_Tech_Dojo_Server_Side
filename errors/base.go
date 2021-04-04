package errors

import (
	"fmt"
	"net/http"

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

var IDValueError = &BaseError{
	statusCode: http.StatusInternalServerError,
	level:      log.Error,
	msg:        "idは1以上の整数である必要があります。",
}

var NameValueError = &BaseError{
	statusCode: http.StatusBadRequest,
	level:      log.Info,
	msg:        "nameは1文字以上である必要があります。",
}

var TokenValueError = &BaseError{
	statusCode: http.StatusInternalServerError,
	level:      log.Error,
	msg:        "tokenは1文字以上である必要があります。",
}
