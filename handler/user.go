package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/application"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/service"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/repository"
)

type createUserJson struct {
	Name string `json:"name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
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

	ur := repository.NewUserRepository()
	uis := service.NewUserIdService(ur)
	uts := service.NewUserTokenService(ur)

	uas := application.NewUserApplicationService(ur, uis, uts)

	token, err := uas.Register(name)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("token", *token)
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	stringToken := header["X-Token"][0] // なんで大文字になる？

	ur := repository.NewUserRepository()
	uis := service.NewUserIdService(ur)
	uts := service.NewUserTokenService(ur)

	uas := application.NewUserApplicationService(ur, uis, uts)

	name, err := uas.GetName(stringToken)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("name", *name)
	w.WriteHeader(http.StatusOK)
}

type updateUserJson struct {
	Name string `json:"name"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
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

	ur := repository.NewUserRepository()
	uis := service.NewUserIdService(ur)
	uts := service.NewUserTokenService(ur)

	uas := application.NewUserApplicationService(ur, uis, uts)

	err := uas.Update(name, token)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}