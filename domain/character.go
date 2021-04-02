package domain

import (
	"golang.org/x/xerrors"
)

type Character struct {
	id     CharacterID
	name   CharacterName
	rarity CharacterRarity
}

func (c *Character) GetID() CharacterID {
	return c.id
}

func (c *Character) GetName() CharacterName {
	return c.name
}

func (c *Character) GetRarity() CharacterRarity {
	return c.rarity
}

func NewCharacter(id CharacterID, name CharacterName, rarity CharacterRarity) Character {
	return Character{id, name, rarity}
}

type CharacterID int
type CharacterName string
type CharacterRarity int

func NewCharacterID(value int) (*CharacterID, error) {
	if value < 1 {
		return nil, xerrors.New("idは1以上の整数である必要があります。")
	}
	id := CharacterID(value)
	return &id, nil
}

func NewCharacterName(value string) (*CharacterName, error) {
	if len(value) < 1 {
		return nil, xerrors.New("nameは1文字以上である必要があります。")
	}
	name := CharacterName(value)
	return &name, nil
}

func NewCharacterRarity(value int) (*CharacterRarity, error) {
	if value < 1 {
		return nil, xerrors.New("rarityは1以上の整数である必要があります。")
	}
	if 100000 <= value {
		return nil, xerrors.New("rarityは100000以下の整数である必要があります。")
	}
	id := CharacterRarity(value)
	return &id, nil
}
