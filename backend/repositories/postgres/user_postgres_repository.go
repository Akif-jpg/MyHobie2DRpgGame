package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Akif-jpg/MyHobieMMORPGGame/models"
	_ "github.com/lib/pq"
)

// UserPostgresRepository implements repositories.UserRepositoryInterface using PostgreSQL.
type UserPostgresRepository struct {
	db *sql.DB
}

// NewUserPostgresRepository creates a new UserPostgresRepository.
func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, email, role, password, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	return user, nil
}

func (r *UserPostgresRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, role, password, created_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return user, nil
}

func (r *UserPostgresRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, role, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.Email, user.Role, user.Password, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("CreateUser: %w", err)
	}
	return nil
}

func (r *UserPostgresRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, role = $3, password = $4 WHERE id = $5`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Role, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}
	return nil
}

func (r *UserPostgresRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteUser: %w", err)
	}
	return nil
}
