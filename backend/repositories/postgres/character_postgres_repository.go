package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	_ "github.com/lib/pq"
)

// CharacterPostgresRepository implements repositories.CharacterRepositoryInterface using PostgreSQL.
type CharacterPostgresRepository struct {
	db *sql.DB
}

// NewCharacterPostgresRepository creates a new CharacterPostgresRepository.
func NewCharacterPostgresRepository(db *sql.DB) *CharacterPostgresRepository {
	return &CharacterPostgresRepository{db: db}
}

func (r *CharacterPostgresRepository) GetCharacterByID(id int) (*models.Character, error) {
	query := `SELECT id, name, level, experience FROM characters WHERE id = $1`
	row := r.db.QueryRow(query, id)

	c := &models.Character{}
	err := row.Scan(&c.ID, &c.Name, &c.Level, &c.Experience)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("character with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("GetCharacterByID: %w", err)
	}
	return c, nil
}

func (r *CharacterPostgresRepository) GetCharactersByUserID(userID int) ([]*models.Character, error) {
	query := `SELECT id, name, level, experience FROM characters WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("GetCharactersByUserID: %w", err)
	}
	defer rows.Close()

	var characters []*models.Character
	for rows.Next() {
		c := &models.Character{}
		if err := rows.Scan(&c.ID, &c.Name, &c.Level, &c.Experience); err != nil {
			return nil, fmt.Errorf("GetCharactersByUserID scan: %w", err)
		}
		characters = append(characters, c)
	}
	return characters, rows.Err()
}

func (r *CharacterPostgresRepository) CreateCharacter(character *models.Character) error {
	query := `INSERT INTO characters (name, level, experience) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, character.Name, character.Level, character.Experience).Scan(&character.ID)
	if err != nil {
		return fmt.Errorf("CreateCharacter: %w", err)
	}
	return nil
}

func (r *CharacterPostgresRepository) UpdateCharacter(character *models.Character) error {
	query := `UPDATE characters SET name = $1, level = $2, experience = $3 WHERE id = $4`
	_, err := r.db.Exec(query, character.Name, character.Level, character.Experience, character.ID)
	if err != nil {
		return fmt.Errorf("UpdateCharacter: %w", err)
	}
	return nil
}

func (r *CharacterPostgresRepository) DeleteCharacter(id int) error {
	query := `DELETE FROM characters WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteCharacter: %w", err)
	}
	return nil
}
