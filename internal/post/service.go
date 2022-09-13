package post

import (
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
)

type post struct{
	repository PostRepository
}

func NewPostService(repository PostRepository) Post {
	return &post{
		repository: repository,
	}
}

func (p *post) Get(id uint64) (domain.Post, error) {
	return p.repository.GetByID(id)
}
func (p *post) GetAll() ([]domain.Post, error) {
	return p.repository.GetAll()
}
func (p *post) GetByUserID(userId uint64) ([]domain.Post, error) {
	return p.repository.GetByUserID(userId)
}
func (p *post) Create(post domain.Post) error {
	return p.repository.Create(post)
}
func (p *post) Update(id uint64, post domain.Post) error {
	return p.repository.Update(id, post)
}
func (p *post) Delete(id uint64) error {
	return p.repository.Delete(id)
}

/*
func (p *post) Get(id uint64) (domain.Post, error) {
	var posts domain.Post
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		return posts, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return posts, err
	}

	if err := json.Unmarshal(body, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *post) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return posts, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return posts, err
	}

	if err := json.Unmarshal(body, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *post) GetAllByUser(id uint64) ([]domain.Post, error) {
	var posts []domain.Post
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d/posts", id))
	if err != nil {
		return posts, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return posts, err
	}

	if err := json.Unmarshal(body, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}
 */
