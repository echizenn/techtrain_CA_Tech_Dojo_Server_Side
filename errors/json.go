package errors

import (
	"net/http"
)

var JsonMarshalError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "jsonのMarshalに失敗しました",
}
