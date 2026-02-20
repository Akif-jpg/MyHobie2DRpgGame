package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// InventoryPostgresRepository implements repositories.InventoryRepositoryInterface using PostgreSQL.
type InventoryPostgresRepository struct {
	db *sql.DB
}

// NewInventoryPostgresRepository creates a new InventoryPostgresRepository.
func NewInventoryPostgresRepository(db *sql.DB) *InventoryPostgresRepository {
	return &InventoryPostgresRepository{db: db}
}

func (r *InventoryPostgresRepository) GetInventoryByID(id int) (*models.Inventory, error) {
	query := `SELECT id, user_id, character_id, item_list, quantity, created_at FROM inventories WHERE id = $1`
	row := r.db.QueryRow(query, id)
	return scanInventory(row)
}

func (r *InventoryPostgresRepository) GetInventoryByUserID(userID int) (*models.Inventory, error) {
	query := `SELECT id, user_id, character_id, item_list, quantity, created_at FROM inventories WHERE user_id = $1`
	row := r.db.QueryRow(query, userID)
	return scanInventory(row)
}

func (r *InventoryPostgresRepository) GetInventoryByCharacterID(characterID int) (*models.Inventory, error) {
	query := `SELECT id, user_id, character_id, item_list, quantity, created_at FROM inventories WHERE character_id = $1`
	row := r.db.QueryRow(query, characterID)
	return scanInventory(row)
}

func (r *InventoryPostgresRepository) CreateInventory(inventory *models.Inventory) error {
	query := `INSERT INTO inventories (user_id, character_id, item_list, quantity, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query,
		inventory.UserID,
		inventory.CharacterID,
		pq.Array(inventory.ItemList),
		inventory.Quantity,
		inventory.CreatedAt,
	).Scan(&inventory.ID)
	if err != nil {
		return fmt.Errorf("CreateInventory: %w", err)
	}
	return nil
}

func (r *InventoryPostgresRepository) UpdateInventory(inventory *models.Inventory) error {
	query := `UPDATE inventories SET user_id = $1, character_id = $2, item_list = $3, quantity = $4 WHERE id = $5`
	_, err := r.db.Exec(query,
		inventory.UserID,
		inventory.CharacterID,
		pq.Array(inventory.ItemList),
		inventory.Quantity,
		inventory.ID,
	)
	if err != nil {
		return fmt.Errorf("UpdateInventory: %w", err)
	}
	return nil
}

func (r *InventoryPostgresRepository) DeleteInventory(id int) error {
	query := `DELETE FROM inventories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteInventory: %w", err)
	}
	return nil
}

// scanInventory is a helper that scans a single row into an Inventory model.
func scanInventory(row *sql.Row) (*models.Inventory, error) {
	inv := &models.Inventory{}
	err := row.Scan(
		&inv.ID,
		&inv.UserID,
		&inv.CharacterID,
		pq.Array(&inv.ItemList),
		&inv.Quantity,
		&inv.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("inventory not found")
	}
	if err != nil {
		return nil, fmt.Errorf("scanInventory: %w", err)
	}
	return inv, nil
}
