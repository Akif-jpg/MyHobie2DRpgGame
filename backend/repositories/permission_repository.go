package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

// PermissionRepositoryInterface defines the contract for permission data access.
type PermissionRepositoryInterface interface {
	GetPermissionByID(id int) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(permission *models.Permission) error
	UpdatePermission(permission *models.Permission) error
	DeletePermission(id int) error
}
