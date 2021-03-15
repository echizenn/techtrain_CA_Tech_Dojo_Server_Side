package infrastructure

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Insert(user *domain.User) error {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	rows, err := ur.db.Prepare("INSERT INTO users VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = rows.Exec(user.GetId(), user.GetName(), user.GetToken())
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(user *domain.User) error {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	rows, err := ur.db.Prepare("UPDATE users SET name=? WHERE id=? AND token=?")
	if err != nil {
		return err
	}

	// この辺型大丈夫なのかよくわからない
	_, err = rows.Exec(user.GetName(), user.GetId(), user.GetToken())
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) FindByToken(userToken *domain.UserToken) (*domain.User, error) {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	var id int
	var name string

	err := ur.db.QueryRow("SELECT id, name FROM users WHERE token=?", userToken).Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	userId, err := domain.NewUserId(id)
	if err != nil {
		return nil, err
	}

	userName, err := domain.NewUserName(name)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(*userId, *userName, *userToken)

	return &user, nil
}

func (ur *userRepository) GetMaxId() (*domain.UserId, error) {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	var id int
	err := ur.db.QueryRow("SELECT MAX(id) FROM users").Scan(&id)
	if err != nil {
		return nil, err
	}

	userId, err := domain.NewUserId(id)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
