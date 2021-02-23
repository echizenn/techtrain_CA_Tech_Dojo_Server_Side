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

	http.HandleFunc("/gacha/draw", gachaDraw)

	// このURL微妙な感じがした
	http.HandleFunc("/character/list", userHoldCharacterList)

	// 8088ポートで起動
	http.ListenAndServe(":8088", nil)
}

type createUserJson struct {
	Name string `json:"name"`
}

type updateUserJson struct {
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

func updateUser(w http.ResponseWriter, r *http.Request) {
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

type gachaDrawJson struct {
	Times int `json:"times"`
}

func gachaDraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("POSTだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	var gdj gachaDrawJson
	json.Unmarshal(buf.Bytes(), &gdj)

	times := gdj.Times

	ur := repository.NewUserRepository()
	cr := repository.NewCharacterRepository()
	ucr := repository.NewUsersCharactersRepository(cr)

	gs := service.NewGachaService(cr)

	gas := application.NewGachaApplicationService(ur, ucr, gs)

	var results []application.GachaDrawResult

	for i := 0; i < times; i++ {
		gachaDrawResult, err := gas.Draw(token)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, *gachaDrawResult)
	}

	stringResults, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("results", string(stringResults))
	w.WriteHeader(http.StatusOK)
}

func userHoldCharacterList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// ここの処理怪しさしかない
		w.Write([]byte("GETだけです。"))
		return
	}

	header := r.Header
	token := header["X-Token"][0] // なんで大文字になる？、0って明示して大丈夫？

	ur := repository.NewUserRepository()
	cr := repository.NewCharacterRepository()
	ucr := repository.NewUsersCharactersRepository(cr)

	ucas := application.NewUsersCharactersApplicationService(ur, ucr)

	userHoldCharacters, err := ucas.Hold(token)
	if err != nil {
		log.Fatal(err)
	}

	stringCharacters, err := json.Marshal(userHoldCharacters)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("characters", string(stringCharacters))
	w.WriteHeader(http.StatusOK)
}
