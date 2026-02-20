package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

// ItemRepositoryInterface defines the contract for item data access.
// Items are identified by UUID as per the model definition.
type ItemRepositoryInterface interface {
	GetItemByUUID(uuid string) (*models.Item, error)
	GetAllItems() ([]*models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(uuid string) error
}
