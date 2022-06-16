package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users/login", userHandler.Login)
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/users/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}
