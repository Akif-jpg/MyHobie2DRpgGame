package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	_ "github.com/lib/pq"
)

// PermissionPostgresRepository implements repositories.PermissionRepositoryInterface using PostgreSQL.
type PermissionPostgresRepository struct {
	db *sql.DB
}

// NewPermissionPostgresRepository creates a new PermissionPostgresRepository.
func NewPermissionPostgresRepository(db *sql.DB) *PermissionPostgresRepository {
	return &PermissionPostgresRepository{db: db}
}

func (r *PermissionPostgresRepository) GetPermissionByID(id int) (*models.Permission, error) {
	query := `SELECT id, name, description, created_at FROM permissions WHERE id = $1`
	row := r.db.QueryRow(query, id)
	return scanPermission(row)
}

func (r *PermissionPostgresRepository) GetPermissionByName(name string) (*models.Permission, error) {
	query := `SELECT id, name, description, created_at FROM permissions WHERE name = $1`
	row := r.db.QueryRow(query, name)
	return scanPermission(row)
}

func (r *PermissionPostgresRepository) GetAllPermissions() ([]*models.Permission, error) {
	query := `SELECT id, name, description, created_at FROM permissions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllPermissions: %w", err)
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		p := &models.Permission{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("GetAllPermissions scan: %w", err)
		}
		permissions = append(permissions, p)
	}
	return permissions, rows.Err()
}

func (r *PermissionPostgresRepository) CreatePermission(permission *models.Permission) error {
	query := `INSERT INTO permissions (name, description, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, permission.Name, permission.Description, permission.CreatedAt).Scan(&permission.ID)
	if err != nil {
		return fmt.Errorf("CreatePermission: %w", err)
	}
	return nil
}

func (r *PermissionPostgresRepository) UpdatePermission(permission *models.Permission) error {
	query := `UPDATE permissions SET name = $1, description = $2 WHERE id = $3`
	_, err := r.db.Exec(query, permission.Name, permission.Description, permission.ID)
	if err != nil {
		return fmt.Errorf("UpdatePermission: %w", err)
	}
	return nil
}

func (r *PermissionPostgresRepository) DeletePermission(id int) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeletePermission: %w", err)
	}
	return nil
}

// scanPermission is a helper that scans a single row into a Permission model.
func scanPermission(row *sql.Row) (*models.Permission, error) {
	p := &models.Permission{}
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("permission not found")
	}
	if err != nil {
		return nil, fmt.Errorf("scanPermission: %w", err)
	}
	return p, nil
}
