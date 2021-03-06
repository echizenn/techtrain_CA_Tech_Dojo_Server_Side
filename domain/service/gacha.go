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
	maxID, err := gs.characterRepository.GetMaxID()
	if err != nil {
		return nil, xerrors.Errorf("characterRepository.GetMaxID func error: %w", err)
	}
	var intCharacterID int = 0
	for intCharacterID == 0 {
		// rand.Intnは[0, n)でランダムな整数返す。[1,n]で返して欲しいので+1
		// ガチャする候補のキャラクターを選んでいる
		intRandomCharacterID := rand.Intn(int(*maxID)) + 1
		randomCharacterID, err := domain.NewCharacterID(intRandomCharacterID)
		if err != nil {
			return nil, xerrors.Errorf("NewCharacterID func error: %w", err)
		}
		randomCharacter, err := gs.characterRepository.FindByID(randomCharacterID)
		// 長期的にはIDに欠番があってもガチャ回るようにしたい
		if err != nil {
			return nil, xerrors.Errorf("characterRepository.FindByID func error: %w", err)
		}
		rarity := randomCharacter.GetRarity()
		intRarity := int(rarity)

		// [0, レア度)でランダムに数字を一つ選ぶ
		// その数字が0ならそのキャラクター獲得とする
		// 0でなければ獲得失敗でガチャを再び繰り返す
		// このfor文は1回のガチャで10**5回とか呼ばれると思われる
		result := rand.Intn(intRarity)
		if result == 0 {
			intCharacterID = intRandomCharacterID
		}
	}
	characterID, err := domain.NewCharacterID(intCharacterID)
	if err != nil {
		return nil, xerrors.Errorf("NewCharacterID func error: %w", err)
	}
	character, err := gs.characterRepository.FindByID(characterID)
	if err != nil {
		return nil, xerrors.Errorf("characterRepository.FindByID func error: %w", err)
	}
	return character, nil
}
