package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var UserIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      log.Error,
	Msg:        "userIDは1以上の整数である必要があります。",
}

var UserNameError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      log.Info,
	Msg:        "userNameは1文字以上である必要があります。",
}

var UserTokenError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      log.Error,
	Msg:        "userTokenは1文字以上である必要があります。",
}
