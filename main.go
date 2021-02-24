package main

import (
	"net/http"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/handler"
)

func main() {
	// これは書き方よくなさそう
	http.HandleFunc("/user/create", handler.CreateUser)
	http.HandleFunc("/user/get", handler.GetUser)
	http.HandleFunc("/user/update", handler.UpdateUser)

	http.HandleFunc("/gacha/draw", handler.GachaDraw)

	// このURL微妙な感じがした
	http.HandleFunc("/character/list", handler.UserHoldCharacterList)

	// 8088ポートで起動
	http.ListenAndServe(":8088", nil)
}
