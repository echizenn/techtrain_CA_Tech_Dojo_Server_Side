package infrastructure

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Insert(user *domain.User) error {
	rows, err := ur.db.Prepare("INSERT INTO users VALUES (?, ?, ?)")
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	_, err = rows.Exec(user.GetID(), user.GetName(), user.GetToken())
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	return nil
}

func (ur *userRepository) Update(user *domain.User) error {
	rows, err := ur.db.Prepare("UPDATE users SET name=? WHERE id=? AND token=?")
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	// この辺型大丈夫なのかよくわからない
	_, err = rows.Exec(user.GetName(), user.GetID(), user.GetToken())
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	return nil
}

func (ur *userRepository) FindByToken(userToken *domain.UserToken) (*domain.User, error) {
	var id int
	var name string

	err := ur.db.QueryRow("SELECT id, name FROM users WHERE token=?", userToken).Scan(&id, &name)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userID, err := domain.NewUserID(id)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userName, err := domain.NewUserName(name)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	user := domain.NewUser(*userID, *userName, *userToken)

	return &user, nil
}

func (ur *userRepository) GetMaxID() (*domain.UserID, error) {
	var id int
	err := ur.db.QueryRow("SELECT MAX(id) FROM users").Scan(&id)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	userID, err := domain.NewUserID(id)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}

	return userID, nil
}
