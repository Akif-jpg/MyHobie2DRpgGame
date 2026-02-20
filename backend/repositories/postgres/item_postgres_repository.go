package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	_ "github.com/lib/pq"
)

// ItemPostgresRepository implements repositories.ItemRepositoryInterface using PostgreSQL.
// Items are keyed by UUID (string).
type ItemPostgresRepository struct {
	db *sql.DB
}

// NewItemPostgresRepository creates a new ItemPostgresRepository.
func NewItemPostgresRepository(db *sql.DB) *ItemPostgresRepository {
	return &ItemPostgresRepository{db: db}
}

func (r *ItemPostgresRepository) GetItemByUUID(uuid string) (*models.Item, error) {
	query := `SELECT id, name, icon, effects, description, price, quantity, created_at FROM items WHERE id = $1`
	row := r.db.QueryRow(query, uuid)

	item := &models.Item{}
	err := row.Scan(&item.UUID, &item.Name, &item.Icon, &item.Effects, &item.Description, &item.Price, &item.Quantity, &item.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item with uuid %s not found", uuid)
	}
	if err != nil {
		return nil, fmt.Errorf("GetItemByUUID: %w", err)
	}
	return item, nil
}

func (r *ItemPostgresRepository) GetAllItems() ([]*models.Item, error) {
	query := `SELECT id, name, icon, effects, description, price, quantity, created_at FROM items`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllItems: %w", err)
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		if err := rows.Scan(&item.UUID, &item.Name, &item.Icon, &item.Effects, &item.Description, &item.Price, &item.Quantity, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("GetAllItems scan: %w", err)
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *ItemPostgresRepository) CreateItem(item *models.Item) error {
	query := `INSERT INTO items (id, name, icon, effects, description, price, quantity, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, item.UUID, item.Name, item.Icon, item.Effects, item.Description, item.Price, item.Quantity, item.CreatedAt)
	if err != nil {
		return fmt.Errorf("CreateItem: %w", err)
	}
	return nil
}

func (r *ItemPostgresRepository) UpdateItem(item *models.Item) error {
	query := `UPDATE items SET name = $1, icon = $2, effects = $3, description = $4, price = $5, quantity = $6 WHERE id = $7`
	_, err := r.db.Exec(query, item.Name, item.Icon, item.Effects, item.Description, item.Price, item.Quantity, item.UUID)
	if err != nil {
		return fmt.Errorf("UpdateItem: %w", err)
	}
	return nil
}

func (r *ItemPostgresRepository) DeleteItem(uuid string) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return fmt.Errorf("DeleteItem: %w", err)
	}
	return nil
}
