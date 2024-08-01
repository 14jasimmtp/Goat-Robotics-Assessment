package handler

import (
	"net/http"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	interfaceUsecase "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase/inteface"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Usecase interfaceUsecase.AuthUsecase
}

func NewAuthHandler(us interfaceUsecase.AuthUsecase) AuthHandler {
	return AuthHandler{Usecase: us}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var body models.Register

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Usecase.Register(body)
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

	user, err := h.Usecase.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
