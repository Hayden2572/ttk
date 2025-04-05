package main

import (
	"log"
	"os"
	"ttk/handlers"
	"ttk/middleware"
	"ttk/repositories"
	"ttk/services"
	"ttk/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Инициализация базы данных
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepo, os.Getenv("JWT_SECRET"))
	userService := services.NewUserService(userRepo)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Создание роутера Gin
	r := gin.Default()

	// Публичные маршруты
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Приватные маршруты (требуют JWT)
	private := r.Group("/")
	private.Use(middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")))
	{
		private.GET("/user/:id", userHandler.GetUser)
		private.PUT("/user/:id", userHandler.UpdateUser)
	}

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
