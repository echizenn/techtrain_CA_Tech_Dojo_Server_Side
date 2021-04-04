package errors

import (
	"net/http"
)

var UserIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "userIDは1以上の整数である必要があります。",
}

var UserNameError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      Info,
	Msg:        "userNameは1文字以上である必要があります。",
}

var UserTokenError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "userTokenは1文字以上である必要があります。",
}
