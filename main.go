package main

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/api"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/api/wire"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/db/mysql"
)

func main() {
	// dbインスタンス作成
	db, err := mysql.CreateSQLInstance()
	defer db.Close()
	if err != nil {
		// dbインスタンスが立ち上がらなかった時の処理を書く
	}

	gameAPI := wire.InitGameAPI(db)

	mux := Server(gameAPI)
	if err := http.ListenAndServe(":8088", mux); err != nil {
		// listen and serveに失敗した時の処理を書く
	}
}

type methodHandler map[string]http.Handler

func (m methodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m[r.Method]; ok {
		h.ServeHTTP(w, r)
		return
	}
	http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
}

func Server(gameAPI api.GameAPI) *http.ServeMux {
	mux := http.NewServeMux()
	// user
	mux.Handle("/user/create", methodHandler{http.MethodPost: http.HandlerFunc(gameAPI.CreateUserHandler)})
	mux.Handle("/user/get", methodHandler{http.MethodGet: http.HandlerFunc(gameAPI.GetUserHandler)})
	mux.Handle("/user/update", methodHandler{http.MethodPut: http.HandlerFunc(gameAPI.UpdateUserHandler)})

	// gacha
	mux.Handle("/gacha/draw", methodHandler{http.MethodPost: http.HandlerFunc(gameAPI.GachaDrawHandler)})

	// character
	mux.Handle("/character/list", methodHandler{http.MethodGet: http.HandlerFunc(gameAPI.UserHoldCharacterListHandler)})
	return mux
}
