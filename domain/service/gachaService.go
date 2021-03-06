package service

import (
	"log"
	"math/rand"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository_interface"
)

type GachaService struct {
	characterRepository repository_interface.ICharacterRepository
}

func NewGachaService(characterRepository repository_interface.ICharacterRepository) GachaService {
	return GachaService{characterRepository}
}

func (gs *GachaService) Draw() *domain.Character {
	maxId, err := gs.characterRepository.GetMaxId()
	if err != nil {
		log.Fatal(err)
	}
	var intCharacterId int = 0
	for intCharacterId == 0 {
		// rand.Intnは[0, n)でランダムな整数返す。[1,n]で返して欲しいので+1
		// ガチャする候補のキャラクターを選んでいる
		intRandomCharacterId := rand.Intn(int(*maxId)) + 1
		randomCharacterId, err := domain.NewCharacterId(intRandomCharacterId)
		if err != nil {
			log.Fatal(err)
		}
		randomCharacter, err := gs.characterRepository.FindById(randomCharacterId)
		// 長期的にはIdに欠番があってもガチャ回るようにしたい
		if err != nil {
			log.Fatal(err)
		}
		rarity := randomCharacter.GetRarity()
		intRarity := int(rarity)

		// [0, レア度)でランダムに数字を一つ選ぶ
		// その数字が0ならそのキャラクター獲得とする
		// 0でなければ獲得失敗でガチャを再び繰り返す
		// このfor文は1回のガチャで10**5回とか呼ばれると思われる
		// sql呼びすぎな気もする
		// 計算量削減したいならあらかじめ全てのキャラクターのレア度から一発でガチャができるようにする
		result := rand.Intn(intRarity)
		if result == 0 {
			intCharacterId = intRandomCharacterId
		}
	}
	characterId, err := domain.NewCharacterId(intCharacterId)
	if err != nil {
		log.Fatal(err)
	}
	character, err := gs.characterRepository.FindById(characterId)
	if err != nil {
		log.Fatal(err)
	}
	return character
}
