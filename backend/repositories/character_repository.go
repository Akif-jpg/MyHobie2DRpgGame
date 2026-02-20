package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

// CharacterRepositoryInterface defines the contract for character data access.
type CharacterRepositoryInterface interface {
	GetCharacterByID(id int) (*models.Character, error)
	GetCharactersByUserID(userID int) ([]*models.Character, error)
	CreateCharacter(character *models.Character) error
	UpdateCharacter(character *models.Character) error
	DeleteCharacter(id int) error
}
