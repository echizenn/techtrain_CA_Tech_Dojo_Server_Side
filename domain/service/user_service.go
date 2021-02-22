package service

import (
	"log"

	"github.com/google/uuid"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository_interface"
)

type UserTokenService struct {
	userRepository repository_interface.IUserRepository
}

func NewUserTokenService(userRepository repository_interface.IUserRepository) UserTokenService {
	return UserTokenService{userRepository}
}

func (uts UserTokenService) Create() *domain.UserToken {
	// ランダム生成
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	stringToken := u.String()
	token, _ := domain.NewToken(stringToken)
	// トークンが衝突している場合は新しいトークン発行
	for uts.Exists(*token) {
		// ランダム生成
		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}
		stringToken := u.String()
		token, _ = domain.NewToken(stringToken)
	}
	return token
}

func (uts UserTokenService) Exists(token domain.UserToken) bool {
	_, err := uts.userRepository.FindByToken(&token)
	if err != nil {
		// この処理いいか微妙
		return false
	}
	return true
}

type UserIdService struct {
	userRepository repository_interface.IUserRepository
}

func NewUserIdService(userRepository repository_interface.IUserRepository) UserIdService {
	return UserIdService{userRepository}
}

func (uis UserIdService) Create() *domain.UserId {
	// 現在最大のidを取得
	maxUserId, err := uis.userRepository.GetMaxId()
	var newId *domain.UserId

	// 登録者いないとき(この方法でハンドリングしていいの？)
	if err != nil {
		newId, err = domain.NewId(1)
		if err != nil {
			log.Fatal(err)
		}
		// 登録者いるとき
	} else {
		id := int(*maxUserId) + 1
		newId, err = domain.NewId(id)
		if err != nil {
			log.Fatal(err)
		}
	}

	return newId
}
