package infrastructure

import (
	"database/sql"

	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain"
	"github.com/echizenn/techtrain_CA_Tech_Dojo_Server_Side/domain/repository"
	_ "github.com/go-sql-driver/mysql"
)

type characterRepository struct {
	db *sql.DB
}

func NewCharacterRepository(db *sql.DB) repository.ICharacterRepository {
	return &characterRepository{db}
}

func (cr *characterRepository) FindById(characterId *domain.CharacterId) (*domain.Character, error) {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	var rarity int
	var name string

	err := cr.db.QueryRow("SELECT name, rarity FROM characters WHERE id=?", characterId).Scan(&name, &rarity)
	if err != nil {
		return nil, err
	}

	characterRarity, err := domain.NewCharacterRarity(rarity)
	if err != nil {
		return nil, err
	}

	characterName, err := domain.NewCharacterName(name)
	if err != nil {
		return nil, err
	}

	character := domain.NewCharacter(*characterId, *characterName, *characterRarity)

	return &character, nil
}

func (cr *characterRepository) GetMaxId() (*domain.CharacterId, error) {
	//db := connectMysql.CreateSQLInstance()
	//defer db.Close()

	var id int
	err := cr.db.QueryRow("SELECT MAX(id) FROM characters").Scan(&id)
	if err != nil {
		return nil, err
	}

	userId, err := domain.NewCharacterId(id)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
