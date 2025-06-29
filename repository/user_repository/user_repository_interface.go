package userrepository

import "github.com/imnzr/northern-gaming-club-api/models"

type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	Update(userId int) error
	GetAllUser() ([]models.User, error)
	GetUserId(userId int) error
}
