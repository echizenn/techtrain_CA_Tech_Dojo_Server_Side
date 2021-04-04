package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"
)

type createUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) CreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけです。"))
		return nil
	}

	var cuj createUserJson
	body := r.Body
	dec := json.NewDecoder(body)
	dec.Decode(&cuj)

	name := cuj.Name

	token, err := api.userApplicationService.Register(name)
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	w.Header().Set("token", *token)

	return nil
}

func (api *GameAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := api.CreateUser(w, r)
	if err != nil {
		// statusコードを設定
		// ログをはく
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *GameAPI) GetUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("GETだけです。"))
		return nil
	}

	header := r.Header
	stringToken := header["X-Token"][0] // なんで大文字になる？

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
		// statusコードを設定
		// ログをはく
		return
	}
	w.WriteHeader(http.StatusOK)
}

type updateUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("PUTだけです。"))
		return nil
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

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
		// statusコードを設定
		// ログをはく
		return
	}
	w.WriteHeader(http.StatusOK)
}
