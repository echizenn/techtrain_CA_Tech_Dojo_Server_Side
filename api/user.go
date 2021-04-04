package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
)

type createUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) CreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.MethodNotAllowedError
	}

	var cuj createUserJson
	body := r.Body
	dec := json.NewDecoder(body)
	dec.Decode(&cuj)

	name := cuj.Name

	token, err := api.userApplicationService.Register(name)
	if err != nil {
		return xerrors.Errorf("userApplicationService.Register func error: %w", err)
	}

	w.Header().Set("token", *token)

	return nil
}

func (api *GameAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := api.CreateUser(w, r)
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

func (api *GameAPI) GetUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.MethodNotAllowedError
	}

	header := r.Header
	stringToken := header.Get("X-Token")
	if stringToken == "" {
		return errors.NoTokenError
	}

	name, err := api.userApplicationService.GetName(stringToken)
	if err != nil {
		return xerrors.Errorf("userApplicationService.GetName func error: %w", err)
	}

	w.Header().Set("name", *name)

	return nil
}

func (api *GameAPI) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	err := api.GetUser(w, r)
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

type updateUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return errors.MethodNotAllowedError
	}

	header := r.Header
	token := header.Get("X-Token")
	if token == "" {
		return errors.NoTokenError
	}

	var uuj updateUserJson
	body := r.Body
	dec := json.NewDecoder(body)
	dec.Decode(&uuj)

	name := uuj.Name

	err := api.userApplicationService.Update(name, token)
	if err != nil {
		return xerrors.Errorf("userApplicationService.Update func error: %w", err)
	}

	return nil
}

func (api *GameAPI) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := api.UpdateUser(w, r)
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
