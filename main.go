package main

import (
	"github.com/gin-gonic/gin"
	"tdd/handler"
	"tdd/repository"
	"tdd/usecase"
)

func main() {
	r := gin.Default()

	// Initialize repository
	repo := repository.NewMemoryRepository()

	// Initialize usecase
	uc := usecase.NewUseCase(repo)

	// Initialize handler
	h := handler.NewHandler(uc)

	// Define routes
	r.POST("/data", h.PostData)

	// Run the server
	r.Run(":8080")
}
