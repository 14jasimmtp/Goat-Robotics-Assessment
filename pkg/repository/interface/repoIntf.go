package interfaceRepo

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
)

type AuthInterface interface {
	SaveUser(user models.Register, Password string) (*db.Users, error)
	FetchUserByEmail(email string) (*db.Users, error)
}
