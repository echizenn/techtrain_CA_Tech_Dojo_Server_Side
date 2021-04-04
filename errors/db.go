package errors

import (
	"net/http"
)

var DBError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "dbでエラーが発生しました",
}

var OpenDBError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "dbの接続に失敗しました",
}
