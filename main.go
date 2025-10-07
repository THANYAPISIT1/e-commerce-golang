package main

import (
	"e-commerce/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	userDelerivery "e-commerce/features/user/delivery"
	userRepository "e-commerce/features/user/repository"
	userUsecase "e-commerce/features/user/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := database.ConnectPostgres(); err != nil {
		panic("failed to connect database")
	}

	if err := database.AutoMigrate(); err != nil {
		panic("failed to auto migrate")
	}

	log.Println("Database connected successfully!")
}

func main() {
	r := gin.Default()

	v1Group := r.Group("/v1")

	userDelerivery.NewHandler(v1Group, userUsecase.NewUserUsecase(
		userRepository.NewUserRepository(database.PostgresDB),
	))

	r.Run()
}
