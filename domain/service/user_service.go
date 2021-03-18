package service

import (
	"log"

	"github.com/google/uuid"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
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
		return nil, err
	}
	stringToken := u.String()
	token, _ := domain.NewUserToken(stringToken)
	// トークンが衝突している場合は新しいトークン発行
	for uts.Exists(*token) {
		// ランダム生成
		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}
		stringToken := u.String()
		token, _ = domain.NewUserToken(stringToken)
	}
	return token, nil
}

func (uts *UserTokenService) Exists(token domain.UserToken) bool {
	_, err := uts.userRepository.FindByToken(&token)
	if err != nil {
		// この処理いいか微妙(よくない、FindByTokenでない場合は""を返す、エラーはエラーで返す)
		return false
	}
	return true
}

type UserIdService struct {
	userRepository repository.IUserRepository
}

func NewUserIdService(userRepository repository.IUserRepository) UserIdService {
	return UserIdService{userRepository}
}

func (uis *UserIdService) Create() (*domain.UserId, error) {
	// 現在最大のidを取得
	maxUserId, err := uis.userRepository.GetMaxId()
	var newId *domain.UserId

	// 登録者いないとき(この方法でハンドリングしていいの？)
	if err != nil {
		newId, err = domain.NewUserId(1)
		if err != nil {
			return nil, err
		}
		// 登録者いるとき
	} else {
		id := int(*maxUserId) + 1
		newId, err = domain.NewUserId(id)
		if err != nil {
			return nil, err
		}
	}

	return newId, nil
}
