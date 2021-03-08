package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/constants"
	_ "github.com/go-sql-driver/mysql"
)

func CreateSQLInstance() *sql.DB {
	dbuser := constants.MysqlDefaultUser
	dbpassword := constants.MysqlDefaultPassword
	dbname := constants.MysqlDefaultName

	dataSource := fmt.Sprintf("%s:%s@/%s", dbuser, dbpassword, dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
