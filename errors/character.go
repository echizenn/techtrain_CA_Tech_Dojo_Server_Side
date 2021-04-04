package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var CharacterIDError = &BaseError{
	StatusCode: http.StatusInternalServerError,
	Level:      log.Error,
	Msg:        "characterIDは1以上の整数である必要があります。",
}

var CharacterNameError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      log.Info,
	Msg:        "characterNameは1文字以上である必要があります。",
}

var CharacterRarityError = &BaseError{
	StatusCode: http.StatusBadRequest,
	Level:      log.Info,
	Msg:        "characterRarityは1以上100000以下の整数である必要があります。",
}
