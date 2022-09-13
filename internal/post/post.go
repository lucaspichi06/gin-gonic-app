package post

import (
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
)

type Post interface {
	Get(uint64) (domain.Post, error)
	GetAll() ([]domain.Post, error)
	GetByUserID(userId uint64) ([]domain.Post, error)
	Create(post domain.Post) error
	Update(id uint64, post domain.Post) error
	Delete(id uint64) error
}
