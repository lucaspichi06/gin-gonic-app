package domain

type Post struct {
	UserID uint64 `json:"userId"`
	ID     uint64 `json:"id"`
	Tittle string `json:"tittle"`
	Body   string `json:"body"`
}
