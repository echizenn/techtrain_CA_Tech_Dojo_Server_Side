package infrastructure

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/errors"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

type characterRepository struct {
	db *sql.DB
}

func NewCharacterRepository(db *sql.DB) repository.ICharacterRepository {
	return &characterRepository{db}
}

func (cr *characterRepository) FindByID(characterID *domain.CharacterID) (*domain.Character, error) {
	var rarity int
	var name string

	err := cr.db.QueryRow("SELECT name, rarity FROM characters WHERE id=?", characterID).Scan(&name, &rarity)
	if err != nil {
		DBError := errors.DBError
		DBError.Msg = err.Error()
		return nil, DBError
	}

	characterRarity, err := domain.NewCharacterRarity(rarity)
	if err != nil {
		return nil, xerrors.Errorf("NewCharacterRarity func error: %w", err)
	}

	characterName, err := domain.NewCharacterName(name)
	if err != nil {
		return nil, xerrors.Errorf("NewCharacterName func error: %w", err)
	}

	character := domain.NewCharacter(*characterID, *characterName, *characterRarity)

	return &character, nil
}

func (cr *characterRepository) GetMaxID() (*domain.CharacterID, error) {
	var id int
	err := cr.db.QueryRow("SELECT MAX(id) FROM characters").Scan(&id)
	if err != nil {
		DBError := errors.DBError
		DBError.Msg = err.Error()
		return nil, DBError
	}

	userID, err := domain.NewCharacterID(id)
	if err != nil {
		return nil, xerrors.Errorf("NewCharacterID func error: %w", err)
	}

	return userID, nil
}
