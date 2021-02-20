package main

import (
    "database/sql"
	"fmt"
    "log"
    "net/http"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    //DBの接続
    var Db *sql.DB
    var err error
    //<user名>:<パスワード>@/<db名>
    Db, err = sql.Open("mysql", "root:example@/go_database")
    if err != nil {
        log.Fatal("DBエラー")
    }

    db.SetConnMaxLifetime(time.Minute * 3)

    fmt.Println("接続できたお( ＾ω＾)")

    // 「/a」に対して処理を追加
    http.HandleFunc("/a", handler)

    // 8088ポートで起動
    http.ListenAndServe(":8088", nil)
}

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World from Go.")
}