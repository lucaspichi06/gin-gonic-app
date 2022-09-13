package handler

import (
	user2 "github.com/lucaspichi06/gin-gonic-app/internal/user"
)

type user struct {
	service user2.User
}

func NewUser(service user2.User) user {
	return user{
		service: service,
	}
}

