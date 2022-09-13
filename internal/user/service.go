package user

import (
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
)

type user struct{
	repository UserRepository
}

func NewUserService(repository UserRepository) User {
	return &user{
		repository: repository,
	}
}

func (u *user) Get(id uint64) (domain.User, error) {
	return u.repository.GetByID(id)
}

func (u *user) GetAll() ([]domain.User, error) {
	return u.repository.GetAll()
}

func (u *user) Create(user domain.User) error {
	return u.repository.Create(user)
}
func (u *user) Update(id uint64, user domain.User) error {
	return u.repository.Update(id, user)
}
func (u *user) Delete(id uint64) error {
	return u.repository.Delete(id)
}

/*
func (p *user) Get(id uint64) (domain.User, error) {
	var user domain.User
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id))
	if err != nil {
		return user, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (p *user) GetAll() ([]domain.User, error) {
	var users []domain.User
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return users, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return users, err
	}

	if err := json.Unmarshal(body, &users); err != nil {
		return users, err
	}

	return users, nil
}
*/
