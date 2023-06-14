package repositories

import (
	"github.com/diegosorrilha/users-api/models"
)

type UserRepository interface {
	Create(user models.User) (id int64, err error)
	GetByID(id int) (user models.User, err error)
	GetAll() (users []models.User, err error)
	Update(user models.User) (int64, error)
	Delete(id int) (int64, error)
}
