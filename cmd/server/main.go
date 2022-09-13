package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucaspichi06/gin-gonic-app/cmd/server/handler"
	"github.com/lucaspichi06/gin-gonic-app/cmd/server/middleware"
	"github.com/lucaspichi06/gin-gonic-app/internal/post"
	"github.com/lucaspichi06/gin-gonic-app/internal/user"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	postRepo := post.NewRepository("./data/posts.json")
	userRepo := user.NewRepository("./data/users.json")

	postService := post.NewPostService(postRepo)
	userService := user.NewUserService(userRepo)

	postHandler := handler.NewPost(postService)
	userHandler := handler.NewUser(userService)


}
