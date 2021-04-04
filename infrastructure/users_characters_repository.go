package infrastructure

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

type usersCharactersRepository struct {
	cr repository.ICharacterRepository
	db *sql.DB
}

func NewUsersCharactersRepository(cr repository.ICharacterRepository, db *sql.DB) repository.IUsersCharactersRepository {
	return &usersCharactersRepository{cr, db}
}

func (ucr *usersCharactersRepository) Insert(user *domain.User, character *domain.Character) error {
	rows, err := ucr.db.Prepare("INSERT INTO users_characters (user_id, character_id) VALUES (?, ?)")
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	_, err = rows.Exec(user.GetID(), character.GetID())
	if err != nil {
		return xerrors.Errorf("error: %w", err)
	}

	return nil
}

func (ucr *usersCharactersRepository) FindByUser(user *domain.User) (*[]*domain.Character, *[]*int, error) {
	userID := user.GetID()

	var intCharacterID int
	var characters []*domain.Character

	var intID int
	var ids []*int

	rows, err := ucr.db.Query("SELECT id, character_id FROM users_characters WHERE user_id=?", userID)
	if err != nil {
		return nil, nil, xerrors.Errorf("error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&intID, &intCharacterID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error: %w", err)
		}
		characterID, err := domain.NewCharacterID(intCharacterID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error: %w", err)
		}
		character, err := ucr.cr.FindByID(characterID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error: %w", err)
		}
		id := int(intID)
		characters = append(characters, character)
		ids = append(ids, &id)
	}

	return &characters, &ids, nil
}
