package errors

import (
	"net/http"
)

var NoTokenError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      Info,
	Msg:        "tokenがありません",
}

var MethodNotAllowedError = &BaseError{
	StatusCode: http.StatusMethodNotAllowed,
	Level:      Info,
	Msg:        "そのmethodは存在しません。",
}
