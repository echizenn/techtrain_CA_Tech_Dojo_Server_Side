package main

import (
	"database/sql"
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/api"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/db/mysql"
)

func main() {
	// dbインスタンス作成
	db := mysql.CreateSQLInstance()
	defer db.Close()

	mux := Router(db)
	if err := http.ListenAndServe(":8088", mux); err != nil {
		panic(err)
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

func Router(db *sql.DB) *http.ServeMux {
	gameAPI := api.NewGameAPI(db)

	mux := http.NewServeMux()
	// user
	mux.Handle("/user/create", methodHandler{http.MethodPost: http.HandlerFunc(gameAPI.CreateUser)})
	mux.Handle("/user/get", methodHandler{http.MethodGet: http.HandlerFunc(gameAPI.GetUser)})
	mux.Handle("/user/update", methodHandler{http.MethodPut: http.HandlerFunc(gameAPI.UpdateUser)})

	// gacha
	mux.Handle("/gacha/draw", methodHandler{http.MethodPost: http.HandlerFunc(gameAPI.GachaDraw)})

	// character
	mux.Handle("/character/list", methodHandler{http.MethodGet: http.HandlerFunc(gameAPI.UserHoldCharacterList)})
	return mux
}
