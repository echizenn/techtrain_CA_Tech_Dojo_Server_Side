package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var JsonMarshalError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      log.Error,
	Msg:        "jsonのMarshalに失敗しました",
}
