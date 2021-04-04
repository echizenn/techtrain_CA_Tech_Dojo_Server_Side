package domain

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
)

type User struct {
	id    UserID
	name  UserName
	token UserToken
}

func (u *User) GetID() UserID {
	return u.id
}

func (u *User) GetName() UserName {
	return u.name
}

func (u *User) SetName(name UserName) *User {
	u.name = name
	return u
}

func (u *User) GetToken() UserToken {
	return u.token
}

func NewUser(id UserID, name UserName, token UserToken) User {
	return User{id, name, token}
}

type UserID int
type UserName string
type UserToken string

func NewUserID(value int) (*UserID, error) {
	if value < 1 {
		return nil, errors.UserIDError
	}
	id := UserID(value)
	return &id, nil
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < 1 {
		return nil, errors.UserNameError
	}
	name := UserName(value)
	return &name, nil
}

func NewUserToken(value string) (*UserToken, error) {
	if len(value) < 1 {
		return nil, errors.UserTokenError
	}
	token := UserToken(value)
	return &token, nil
}
