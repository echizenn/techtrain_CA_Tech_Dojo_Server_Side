package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var UUIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      log.Error,
	Msg:        "uuidでエラーが発生しました",
}
