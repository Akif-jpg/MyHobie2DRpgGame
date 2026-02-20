package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

// RoleRepositoryInterface defines the contract for role data access.
type RoleRepositoryInterface interface {
	GetRoleByID(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(id int) error
	AddPermissionToRole(roleID int, permissionID int) error
	RemovePermissionFromRole(roleID int, permissionID int) error
}
