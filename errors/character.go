package errors

import (
	"net/http"
)

var CharacterIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      Error,
	Msg:        "characterIDは1以上の整数である必要があります。",
}

var CharacterNameError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      Info,
	Msg:        "characterNameは1文字以上である必要があります。",
}

var CharacterRarityError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      Info,
	Msg:        "characterRarityは1以上100000以下の整数である必要があります。",
}
