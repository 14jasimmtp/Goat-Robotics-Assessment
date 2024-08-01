package main

import (
	"log"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/config"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/routes"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.LoadConfig() != nil {
		log.Fatal("error loading config files")
	}
	DB := db.ConnectToDB()
	AuthRepo := repository.NewAuthRepo(DB)
	AuthUsecase := usecase.NewAuthUsecase(AuthRepo)
	AuthHandler := handler.NewAuthHandler(AuthUsecase)
	ProjectRepo := repository.NewProjectRepo(DB)
	ProjectUsecase := usecase.NewProjectUsecase(ProjectRepo)
	ProjectHandler := handler.NewProjectHandler(ProjectUsecase)

	r := gin.Default()
	routes.Routes(r, AuthHandler, ProjectHandler)

	r.Run(":3000")
}
