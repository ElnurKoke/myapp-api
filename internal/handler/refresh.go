package handler

import (
	"elestial/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) refresh(c *gin.Context) {
	var input model.TokenPair

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid input",
			"details": err.Error(),
		})
		return
	}

	claims, err := h.Service.ParseToken(input.Refresh, []byte(h.Cfg.JWT.RefreshSecret))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	newAccess, err := h.Service.GenerateAccessToken(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access": newAccess})
}
