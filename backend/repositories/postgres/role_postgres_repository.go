package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// RolePostgresRepository implements repositories.RoleRepositoryInterface using PostgreSQL.
type RolePostgresRepository struct {
	db *sql.DB
}

// NewRolePostgresRepository creates a new RolePostgresRepository.
func NewRolePostgresRepository(db *sql.DB) *RolePostgresRepository {
	return &RolePostgresRepository{db: db}
}

func (r *RolePostgresRepository) GetRoleByID(id int) (*models.Role, error) {
	query := `SELECT id, name, permissions, created_at FROM roles WHERE id = $1`
	row := r.db.QueryRow(query, id)
	return scanRole(row)
}

func (r *RolePostgresRepository) GetRoleByName(name string) (*models.Role, error) {
	query := `SELECT id, name, permissions, created_at FROM roles WHERE name = $1`
	row := r.db.QueryRow(query, name)
	return scanRole(row)
}

func (r *RolePostgresRepository) GetAllRoles() ([]*models.Role, error) {
	query := `SELECT id, name, permissions, created_at FROM roles`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllRoles: %w", err)
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.ID, &role.Name, pq.Array(&role.Permissions), &role.CreatedAt); err != nil {
			return nil, fmt.Errorf("GetAllRoles scan: %w", err)
		}
		roles = append(roles, role)
	}
	return roles, rows.Err()
}

func (r *RolePostgresRepository) CreateRole(role *models.Role) error {
	query := `INSERT INTO roles (name, permissions, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, role.Name, pq.Array(role.Permissions), role.CreatedAt).Scan(&role.ID)
	if err != nil {
		return fmt.Errorf("CreateRole: %w", err)
	}
	return nil
}

func (r *RolePostgresRepository) UpdateRole(role *models.Role) error {
	query := `UPDATE roles SET name = $1, permissions = $2 WHERE id = $3`
	_, err := r.db.Exec(query, role.Name, pq.Array(role.Permissions), role.ID)
	if err != nil {
		return fmt.Errorf("UpdateRole: %w", err)
	}
	return nil
}

func (r *RolePostgresRepository) DeleteRole(id int) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteRole: %w", err)
	}
	return nil
}

// AddPermissionToRole adds a permission ID to the role's permissions array.
func (r *RolePostgresRepository) AddPermissionToRole(roleID int, permissionID int) error {
	query := `UPDATE roles SET permissions = array_append(permissions, $1) WHERE id = $2 AND NOT ($1 = ANY(permissions))`
	_, err := r.db.Exec(query, permissionID, roleID)
	if err != nil {
		return fmt.Errorf("AddPermissionToRole: %w", err)
	}
	return nil
}

// RemovePermissionFromRole removes a permission ID from the role's permissions array.
func (r *RolePostgresRepository) RemovePermissionFromRole(roleID int, permissionID int) error {
	query := `UPDATE roles SET permissions = array_remove(permissions, $1) WHERE id = $2`
	_, err := r.db.Exec(query, permissionID, roleID)
	if err != nil {
		return fmt.Errorf("RemovePermissionFromRole: %w", err)
	}
	return nil
}

// scanRole is a helper that scans a single row into a Role model.
func scanRole(row *sql.Row) (*models.Role, error) {
	role := &models.Role{}
	err := row.Scan(&role.ID, &role.Name, pq.Array(&role.Permissions), &role.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("role not found")
	}
	if err != nil {
		return nil, fmt.Errorf("scanRole: %w", err)
	}
	return role, nil
}
