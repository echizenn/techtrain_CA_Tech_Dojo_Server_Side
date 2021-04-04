package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var NoTokenError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      log.Info,
	Msg:        "tokenがありません",
}
