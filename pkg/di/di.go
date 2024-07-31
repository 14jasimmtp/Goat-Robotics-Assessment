package di

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase"
)

func InjectDependencies()(handler.AuthHandler){
	DB := db.ConnectToDB()
	AuthRepo := repository.NewAuthRepo(DB)
	AuthUsecase := usecase.NewAuthUsecase(AuthRepo)
	AuthHandler := handler.NewAuthHandler(AuthUsecase)

	return AuthHandler

}