package main

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

func main() {
	// これは書き方よくなさそう
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)

	// 8088ポートで起動
	http.ListenAndServe(":8088", nil)
}

type createUserJson struct {
	Name string `json:"name"`
}

// リクエストを処理する関数
func createUser(w http.ResponseWriter, r *http.Request) {
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

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	stringToken := header["X-Token"] // なんで大文字になる？

	ur := repository.NewUserRepository()
	uis := service.NewUserIdService(ur)
	uts := service.NewUserTokenService(ur)

	uas := application.NewUserApplicationService(ur, uis, uts)

	// 0と明示していいのかな
	name, err := uas.GetName(stringToken[0])
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("name", *name)
	w.WriteHeader(http.StatusOK)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("PUTだけです。"))
		return
	}
}
