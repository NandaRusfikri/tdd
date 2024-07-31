package handler

import (
	"net/http"
	"tdd/models"
	"tdd/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) PostData(c *gin.Context) {
	var data models.Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateData(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
