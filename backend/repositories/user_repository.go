package repositories

import "github.com/Akif-jpg/MyHobieMMORPGGame/models"

type UserRepositoryInterface interface {
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}
