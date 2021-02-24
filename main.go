package main

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/handler"
)

func main() {
	mux := Router()
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

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	// user
	mux.Handle("/user/create", methodHandler{http.MethodPost: http.HandlerFunc(handler.CreateUser)})
	mux.Handle("/user/get", methodHandler{http.MethodGet: http.HandlerFunc(handler.GetUser)})
	mux.Handle("/user/update", methodHandler{http.MethodPut: http.HandlerFunc(handler.UpdateUser)})

	// gacha
	mux.Handle("/gacha/draw", methodHandler{http.MethodPost: http.HandlerFunc(handler.GachaDraw)})

	// character
	mux.Handle("/character/list", methodHandler{http.MethodGet: http.HandlerFunc(handler.UserHoldCharacterList)})
	return mux
}
