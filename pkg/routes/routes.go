package routes

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler handler.AuthHandler, projectHandler *handler.ProjectHandler) {
	r.POST("/register",authHandler.Register)
	r.POST("/login",authHandler.Login)

	r.POST("/projects",middleware.AuthMiddleware,projectHandler.CreateProject)
	r.POST("/projects/task",middleware.AuthMiddleware,projectHandler.AddTask)
	r.POST("/projects/")

}
