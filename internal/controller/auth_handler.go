package controller

import (
	"filmoteka/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(s *service.UserService) *AuthHandler {
	return &AuthHandler{userService: s}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
