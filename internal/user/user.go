package user

import (
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
)

type User interface {
	Get(id uint64) (domain.User, error)
	GetAll() ([]domain.User, error)
	Create(user domain.User) error
	Update(id uint64, user domain.User) error
	Delete(id uint64) error
}
