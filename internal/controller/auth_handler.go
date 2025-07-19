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

// Login юзеров
// @Summary      Авторизация
// @Description  Получение JWT токена по логину и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      loginRequest  true  "Данные пользователя"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /login [post]
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
