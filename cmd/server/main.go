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

	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.Get)
		users.GET("/:id/posts", postHandler.GetAllByUser)
		users.POST("", middleware.Authorization, userHandler.Create)
		users.PUT("/:id", middleware.Authorization, userHandler.Update)
		users.DELETE("/:id", middleware.Authorization, userHandler.Delete)
	}

	posts := r.Group("/posts")
	{
		posts.GET("", postHandler.GetAll)
		posts.GET("/:id", postHandler.Get)
		posts.POST("", middleware.Authorization, postHandler.Create)
		posts.PUT("/:id", middleware.Authorization, postHandler.Update)
		posts.DELETE("/:id", middleware.Authorization, postHandler.Delete)
	}

	r.Run(":8080")
}
