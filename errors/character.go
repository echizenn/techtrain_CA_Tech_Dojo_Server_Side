package errors

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/log"
)

var CharacterIDError = &BaseError{
	statusCode: http.StatusInternalServerError,
	level:      log.Error,
	msg:        "characterIDは1以上の整数である必要があります。",
}

var CharacterNameError = &BaseError{
	statusCode: http.StatusBadRequest,
	level:      log.Info,
	msg:        "characterNameは1文字以上である必要があります。",
}

var CharacterRarityError = &BaseError{
	statusCode: http.StatusBadRequest,
	level:      log.Info,
	msg:        "characterRarityは1以上100000以下の整数である必要があります。",
}
