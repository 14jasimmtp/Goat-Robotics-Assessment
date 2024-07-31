package routes

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler handler.AuthHandler) {
	r.POST("/register",authHandler.Register)
	r.POST("/login",authHandler.Login)

	
}
