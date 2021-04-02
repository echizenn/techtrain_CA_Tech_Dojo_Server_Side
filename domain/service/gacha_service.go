package service

import (
	"math/rand"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"golang.org/x/xerrors"
)

type GachaService struct {
	characterRepository repository.ICharacterRepository
}

func NewGachaService(characterRepository repository.ICharacterRepository) GachaService {
	return GachaService{characterRepository}
}

func (gs *GachaService) Draw() (*domain.Character, error) {
	maxId, err := gs.characterRepository.GetMaxId()
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}
	var intCharacterID int = 0
	for intCharacterID == 0 {
		// rand.Intnは[0, n)でランダムな整数返す。[1,n]で返して欲しいので+1
		// ガチャする候補のキャラクターを選んでいる
		intRandomCharacterID := rand.Intn(int(*maxId)) + 1
		randomCharacterID, err := domain.NewCharacterID(intRandomCharacterID)
		if err != nil {
			return nil, xerrors.Errorf("error: %w", err)
		}
		randomCharacter, err := gs.characterRepository.FindById(randomCharacterID)
		// 長期的にはIdに欠番があってもガチャ回るようにしたい
		if err != nil {
			return nil, xerrors.Errorf("error: %w", err)
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
			intCharacterID = intRandomCharacterID
		}
	}
	characterID, err := domain.NewCharacterID(intCharacterID)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}
	character, err := gs.characterRepository.FindById(characterID)
	if err != nil {
		return nil, xerrors.Errorf("error: %w", err)
	}
	return character, nil
}
