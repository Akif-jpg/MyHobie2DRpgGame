package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

// InventoryRepositoryInterface defines the contract for inventory data access.
type InventoryRepositoryInterface interface {
	GetInventoryByID(id int) (*models.Inventory, error)
	GetInventoryByUserID(userID int) (*models.Inventory, error)
	GetInventoryByCharacterID(characterID int) (*models.Inventory, error)
	CreateInventory(inventory *models.Inventory) error
	UpdateInventory(inventory *models.Inventory) error
	DeleteInventory(id int) error
}
