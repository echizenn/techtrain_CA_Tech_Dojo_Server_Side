package domain

import (
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
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
		return nil, errors.CharacterIDError
	}
	id := CharacterID(value)
	return &id, nil
}

func NewCharacterName(value string) (*CharacterName, error) {
	if len(value) < 1 {
		return nil, errors.CharacterNameError
	}
	name := CharacterName(value)
	return &name, nil
}

func NewCharacterRarity(value int) (*CharacterRarity, error) {
	if value < 1 {
		return nil, errors.CharacterRarityError
	}
	if 100000 <= value {
		return nil, errors.CharacterRarityError
	}
	id := CharacterRarity(value)
	return &id, nil
}
