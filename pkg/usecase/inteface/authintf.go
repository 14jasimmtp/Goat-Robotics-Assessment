package interfaceUsecase

import "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"

type AuthUsecase interface {
	Register(user models.Register) (*models.RegisterRes, error)
	Login(body models.Login) (*models.LoginRes, error)
}
