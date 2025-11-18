package handler

import (
	"elestial/internal/apperror"
	"elestial/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {

	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid input",
			"details": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	tokens, err := h.Service.Login(ctx, input)
	if err != nil {
		switch err {
		case apperror.ErrWrongPassword, apperror.ErrUserNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"access":  tokens.Access,
		"refresh": tokens.Refresh,
		"expired": 900,
	})

}
