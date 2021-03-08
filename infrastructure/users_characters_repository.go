package infrastructure

import (
	connectMysql "github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/db/mysql"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	_ "github.com/go-sql-driver/mysql"
)

type usersCharactersRepository struct {
	cr repository.ICharacterRepository
}

func NewUsersCharactersRepository(cr repository.ICharacterRepository) repository.IUsersCharactersRepository {
	return &usersCharactersRepository{cr}
}

func (ucr *usersCharactersRepository) Insert(user *domain.User, character *domain.Character) error {
	db := connectMysql.CreateSQLInstance()
	defer db.Close()

	rows, err := db.Prepare("INSERT INTO users_characters (user_id, character_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = rows.Exec(user.GetId(), character.GetId())
	if err != nil {
		return err
	}

	return nil
}

func (ucr *usersCharactersRepository) FindByUser(user *domain.User) (*[]*domain.Character, error) {
	db := connectMysql.CreateSQLInstance()
	defer db.Close()

	userId := user.GetId()

	var intCharacterId int
	var characters []*domain.Character

	rows, err := db.Query("SELECT character_id FROM users_characters WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&intCharacterId)
		if err != nil {
			return nil, err
		}
		characterId, err := domain.NewCharacterId(intCharacterId)
		if err != nil {
			return nil, err
		}
		character, err := ucr.cr.FindById(characterId)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return &characters, nil
}
