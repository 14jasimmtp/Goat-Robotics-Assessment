package handler

import (
	"net/http"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	us usecase.AuthUsecase
}

func NewAuthHandler(us usecase.AuthUsecase) AuthHandler {
	return AuthHandler{us: us}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var body models.Register

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.us.Register(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *AuthHandler) Login(c *gin.Context) {
	var body models.Login

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.us.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
