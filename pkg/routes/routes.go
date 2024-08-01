package routes

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authHandler handler.AuthHandler, projectHandler *handler.ProjectHandler) {
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.POST("/projects", middleware.AuthMiddleware,projectHandler.CreateProject)
	r.GET("/projects", middleware.AuthMiddleware, projectHandler.ListProjects)

	// Task routes
	r.POST("/tasks",middleware.AuthMiddleware, projectHandler.AddTask)
	r.DELETE("/tasks/:id", middleware.AuthMiddleware, projectHandler.DeleteTask)

	// Time Entry routes
	r.POST("/time-entries",middleware.AuthMiddleware, projectHandler.AddTimeEntry)
	r.PUT("/time-entries/:id",middleware.AuthMiddleware, projectHandler.UpdateTimeEntry)
	r.DELETE("/time-entries/:id",middleware.AuthMiddleware, projectHandler.DeleteTimeEntry)
	// r.GET("/time-entries/:id", projectHandler.GetTimeEntry)

}
