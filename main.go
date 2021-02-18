package main

import (
	"fmt"
    "net/http"
)

func main() {

    // 「/a」に対して処理を追加
    http.HandleFunc("/a", handler)

    // 8080ポートで起動
    http.ListenAndServe(":8080", nil)
}

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World from Go.")
}