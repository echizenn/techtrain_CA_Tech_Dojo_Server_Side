package errors

import (
	"net/http"
)

var UUIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "uuidでエラーが発生しました",
}
