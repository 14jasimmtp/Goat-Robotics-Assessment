package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	interfaceRepo "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository/interface"
	interfaceUsecase "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase/inteface"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/utils"
)

type AuthUsecase struct {
	Repo interfaceRepo.AuthInterface
}

func NewAuthUsecase(repo interfaceRepo.AuthInterface) interfaceUsecase.AuthUsecase {
	return &AuthUsecase{Repo: repo}
}

func (u *AuthUsecase) Register(user models.Register) (*models.RegisterRes, error) {
	users, err := u.Repo.FetchUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if users.Email != "" {
		return nil, errors.New(`user already exist`)
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, errors.New(`error while hashing password`)
	}

	SavedUser, err := u.Repo.SaveUser(user, string(HashedPassword))
	if err != nil {
		return nil, err
	}

	return &models.RegisterRes{Status: "success", User: *SavedUser}, nil
}

func (u *AuthUsecase) Login(body models.Login) (*models.LoginRes, error) {
	user, err := u.Repo.FetchUserByEmail(body.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, errors.New(`password wrong`)
	}

	token, err := utils.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &models.LoginRes{Status: "Success", Token: token, User: *user}, nil
}
