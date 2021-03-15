package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type createUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけです。"))
		return
	}

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var cuj createUserJson
	json.Unmarshal(buf.Bytes(), &cuj)

	name := cuj.Name

	token, err := api.uas.Register(name)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("token", *token)
	w.WriteHeader(http.StatusOK)
}

func (api *GameAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	stringToken := header["X-Token"][0] // なんで大文字になる？

	name, err := api.uas.GetName(stringToken)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("name", *name)
	w.WriteHeader(http.StatusOK)
}

type updateUserJson struct {
	Name string `json:"name"`
}

func (api *GameAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("PUTだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var uuj updateUserJson
	json.Unmarshal(buf.Bytes(), &uuj)

	name := uuj.Name

	err := api.uas.Update(name, token)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
