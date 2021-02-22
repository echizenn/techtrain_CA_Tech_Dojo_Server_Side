package repository

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository_interface"
	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct{}

func NewUserRepository() repository_interface.IUserRepository {
	return &userRepository{}
}

func (ur userRepository) Save(user *domain.User) error {
	//DBの接続
	//<user名>:<パスワード>@/<db名>
	db, err := sql.Open("mysql", "root:example@/go_database")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Prepare("INSERT INTO user VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = rows.Exec(user.GetId(), user.GetName(), user.GetToken())
	if err != nil {
		return err
	}

	return nil
}

func (ur userRepository) FindByToken(token *domain.UserToken) (*domain.User, error) {
	//DBの接続
	//<user名>:<パスワード>@/<db名>
	db, err := sql.Open("mysql", "root:example@/go_database")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var id int
	var name string

	err = db.QueryRow("SELECT id, name FROM user WHERE token=?", token).Scan(id, name)
	if err != nil {
		return nil, err
	}
	userId, err := domain.NewId(id)
	if err != nil {
		return nil, err
	}

	userName, err := domain.NewName(name)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(*userId, *userName, *token)

	return &user, nil
}

func (ur userRepository) GetMaxId() (*domain.UserId, error) {
	//DBの接続
	//<user名>:<パスワード>@/<db名>
	db, err := sql.Open("mysql", "root:example@/go_database")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var id int

	err = db.QueryRow("SELECT MAX(id) FROM user").Scan(id)
	if err != nil {
		return nil, err
	}
	userId, err := domain.NewId(id)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
