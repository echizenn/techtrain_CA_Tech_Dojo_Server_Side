package mysql

import (
	"database/sql"
	"fmt"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/constants"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
	_ "github.com/go-sql-driver/mysql"
)

func CreateSQLInstance() (*sql.DB, error) {
	dbuser := constants.MysqlDefaultUser
	dbpassword := constants.MysqlDefaultPassword
	protocal := constants.MysqlDefaultProtocal
	dbname := constants.MysqlDefaultName

	dataSource := fmt.Sprintf("%s:%s@%s/%s", dbuser, dbpassword, protocal, dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, errors.OpenDBError
	}
	return db, nil
}
