package service

import (
	"github.com/google/uuid"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
	"golang.org/x/xerrors"
)

type UserTokenService struct {
	userRepository repository.IUserRepository
}

func NewUserTokenService(userRepository repository.IUserRepository) UserTokenService {
	return UserTokenService{userRepository}
}

func (uts *UserTokenService) Create() (*domain.UserToken, error) {
	// ランダム生成
	u, err := uuid.NewRandom()
	if err != nil {
		UUIDError := errors.UUIDError
		UUIDError.Msg = err.Error()
		return nil, UUIDError
	}
	stringToken := u.String()
	token, err := domain.NewUserToken(stringToken)
	if err != nil {
		return nil, xerrors.Errorf("NewUserToken func error: %w", err)
	}

	// トークンが衝突している場合は新しいトークン発行
	for uts.Exists(*token) {
		// ランダム生成
		u, err := uuid.NewRandom()
		if err != nil {
			UUIDError := errors.UUIDError
			UUIDError.Msg = err.Error()
			return nil, UUIDError
		}
		stringToken := u.String()
		token, err = domain.NewUserToken(stringToken)
		if err != nil {
			return nil, xerrors.Errorf("NewUserToken func error: %w", err)
		}
	}
	return token, nil
}

func (uts *UserTokenService) Exists(token domain.UserToken) bool {
	_, err := uts.userRepository.FindByToken(&token)
	if err != nil {
		// この処理いいか微妙(よくない、FindByTokenで存在しない場合は""を返す、エラーはエラーで返す)
		return false
	}
	return true
}

type UserIDService struct {
	userRepository repository.IUserRepository
}

func NewUserIDService(userRepository repository.IUserRepository) UserIDService {
	return UserIDService{userRepository}
}

func (uis *UserIDService) Create() (*domain.UserID, error) {
	// 現在最大のidを取得
	maxUserID, err := uis.userRepository.GetMaxID()
	var newID *domain.UserID

	// 登録者いないとき(この方法でハンドリングしていいの？)
	if err != nil {
		newID, err = domain.NewUserID(1)
		if err != nil {
			return nil, xerrors.Errorf("NewUserID func error: %w", err)
		}
		// 登録者いるとき
	} else {
		id := int(*maxUserID) + 1
		newID, err = domain.NewUserID(id)
		if err != nil {
			return nil, xerrors.Errorf("NewUserID func error: %w", err)
		}
	}

	return newID, nil
}
