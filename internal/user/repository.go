package user

import (
	"encoding/json"
	"errors"
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
	"os"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(id uint64) (domain.User, error)
	Create(user domain.User) error
	Update(id uint64, user domain.User) error
	Delete(id uint64) error
}

type repository struct {
	path string
}

// NewRepository crea un nuevo repositorio
func NewRepository(path string) UserRepository {
	return &repository{path}
}

// GetAll devuelve todos los users
func (r *repository) GetAll() ([]domain.User, error) {
	return r.loadUsers()
}

// GetByID busca un user por su id
func (r *repository) GetByID(id uint64) (domain.User, error) {
	users, err := r.loadUsers()
	if err != nil {
		return domain.User{}, err
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, errors.New("user not found")

}

// Create agrega un nuevo user
func (r *repository) Create(user domain.User) error {
	users, err := r.loadUsers()
	if err != nil {
		return err
	}
	user.ID = uint64(len(users) + 1)
	users = append(users, user)
	return r.saveUsers(users)
}

// Delete elimina un user
func (r *repository) Delete(id uint64) error {
	users, err := r.loadUsers()
	if err != nil {
		return err
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return r.saveUsers(users)
		}
	}
	return errors.New("user not found")
}

// Update actualiza un user
func (r *repository) Update(id uint64, user domain.User) error {
	users, err := r.loadUsers()
	if err != nil {
		return err
	}
	for i, u := range users {
		if u.ID == id {
			user.ID = id
			users[i] = user
			return r.saveUsers(users)
		}
	}
	return errors.New("user not found")
}

// loadUsers carga los users desde un archivo json
func (r *repository) loadUsers() ([]domain.User, error) {
	file, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// saveUsers guarda los users en un archivo json
func (r *repository) saveUsers(users []domain.User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, bytes, 0644)
}