package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var UserIDError = &BaseError{
	statusCode: http.StatusInternalServerError,
	level:      log.Error,
	msg:        "userIDは1以上の整数である必要があります。",
}

var UserNameError = &BaseError{
	statusCode: http.StatusBadRequest,
	level:      log.Info,
	msg:        "userNameは1文字以上である必要があります。",
}

var UserTokenError = &BaseError{
	statusCode: http.StatusInternalServerError,
	level:      log.Error,
	msg:        "userTokenは1文字以上である必要があります。",
}
