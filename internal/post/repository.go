package post

import (
	"encoding/json"
	"errors"
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
	"os"
)

type PostRepository interface {
	GetAll() ([]domain.Post, error)
	GetByID(id uint64) (domain.Post, error)
	GetByUserID(userId uint64) ([]domain.Post, error)
	Create(post domain.Post) error
	Update(id uint64, post domain.Post) error
	Delete(id uint64) error
}

type repository struct {
	path string
}

// NewRepository crea un nuevo repositorio
func NewRepository(path string) PostRepository {
	return &repository{path}
}

// GetAll devuelve todos los posts
func (r *repository) GetByID(id uint64) (domain.Post, error) {
	posts, err :=  r.loadPosts()
	if err != nil {
		return domain.Post{}, err
	}

	for _, p := range posts {
		if p.ID == id {
			return p, nil
		}
	}

	return domain.Post{}, errors.New("post not found")
}

// GetAll devuelve todos los posts
func (r *repository) GetAll() ([]domain.Post, error) {
	return r.loadPosts()
}

// GetByUserID busca todos los posts asociados a un user id
func (r *repository) GetByUserID(userId uint64) ([]domain.Post, error) {
	posts, err := r.loadPosts()
	if err != nil {
		return nil, err
	}

	var response []domain.Post
	for _, post := range posts {
		if post.UserID == userId {
			response = append(response, post)
		}
	}

	return response, nil
}

// Create agrega un nuevo post
func (r *repository) Create(post domain.Post) error {
	posts, err := r.loadPosts()
	if err != nil {
		return err
	}
	post.ID = uint64(len(posts) + 1)
	posts = append(posts, post)
	return r.savePosts(posts)
}

// Delete elimina un post
func (r *repository) Delete(id uint64) error {
	posts, err := r.loadPosts()
	if err != nil {
		return err
	}
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return r.savePosts(posts)
		}
	}
	return errors.New("post not found")
}

// Update actualiza un post
func (r *repository) Update(id uint64, post domain.Post) error {
	posts, err := r.loadPosts()
	if err != nil {
		return err
	}
	for i, u := range posts {
		if u.ID == id {
			post.ID = id
			posts[i] = post
			return r.savePosts(posts)
		}
	}
	return errors.New("post not found")
}

// loadPosts carga los posts desde un archivo json
func (r *repository) loadPosts() ([]domain.Post, error) {
	file, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var posts []domain.Post
	err = json.Unmarshal(file, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// savePosts guarda los posts en un archivo json
func (r *repository) savePosts(posts []domain.Post) error {
	bytes, err := json.Marshal(posts)
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, bytes, 0644)
}