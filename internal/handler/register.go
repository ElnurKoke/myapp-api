package handler

import (
	"elestial/internal/apperror"
	"elestial/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input model.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid input",
			"details": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err := h.Service.Register(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrUserExists):
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return

		case errors.Is(err, apperror.ErrPasswordDontMatch):
			c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
			return

		case errors.Is(err, apperror.ErrInvalidEmail):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			return

		case errors.Is(err, apperror.ErrInvalidUserName):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username"})
			return

		case errors.Is(err, apperror.ErrShortPassword):
			c.JSON(http.StatusBadRequest, gin.H{"error": "password too weak"})
			return

		case errors.Is(err, apperror.ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusCreated)
}
