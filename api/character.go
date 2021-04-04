package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
)

func (api *GameAPI) UserHoldCharacterList(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.MethodNotAllowedError
	}

	header := r.Header
	token := header.Get("X-Token")
	if token == "" {
		return errors.NoTokenError
	}

	userHoldCharacters, err := api.usersCharactersApplicationService.Hold(token)
	if err != nil {
		return xerrors.Errorf("usersCharactersApplicationService.Hold func error: %w", err)
	}

	marshalCharacters, err := json.Marshal(userHoldCharacters)
	if err != nil {
		return errors.JsonMarshalError
	}

	w.Header().Set("characters", string(marshalCharacters))

	return nil
}

func (api *GameAPI) UserHoldCharacterListHandler(w http.ResponseWriter, r *http.Request) {
	err := api.UserHoldCharacterList(w, r)
	if err != nil {
		errors.EmitLog(err)

		var baseError *errors.BaseError
		if xerrors.As(err, &baseError) {
			w.WriteHeader(baseError.StatusCode)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
