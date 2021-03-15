package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type GameAPI struct {
	db *sql.DB
}

func NewGameAPI(db *sql.DB) GameAPI {
	return GameAPI{db}
}
