package domain

import "errors"

type User struct {
	id    UserId
	name  UserName
	token UserToken
}

func (u *User) GetId() UserId {
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

func NewUser(id UserId, name UserName, token UserToken) User {
	return User{id, name, token}
}

type UserId int
type UserName string
type UserToken string

func NewUserId(value int) (*UserId, error) {
	if value < 1 {
		return nil, errors.New("idは1以上の整数である必要があります。")
	}
	id := UserId(value)
	return &id, nil
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < 1 {
		return nil, errors.New(("nameは1文字以上である必要があります。"))
	}
	name := UserName(value)
	return &name, nil
}

func NewUserToken(value string) (*UserToken, error) {
	if len(value) < 1 {
		return nil, errors.New(("tokenは1文字以上である必要があります。"))
	}
	token := UserToken(value)
	return &token, nil
}
