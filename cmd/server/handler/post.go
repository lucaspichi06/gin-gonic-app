package handler

import (
	post2 "github.com/lucaspichi06/gin-gonic-app/internal/post"
)

type post struct {
	service post2.Post
}

func NewPost(service post2.Post) post {
	return post{
		service: service,
	}
}
