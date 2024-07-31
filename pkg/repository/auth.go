package repository

import (
	"errors"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"gorm.io/gorm"
)

type AuthRepo struct{
	Db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo{
	return AuthRepo{Db: db}
}

func (r *AuthRepo) SaveUser(user models.Register, Password string) (*db.Users, error){
	users:=db.Users{Username: user.Username,Email: user.Email, Password: Password}
	tx:=r.Db.Create(&users)
	if tx.Error != nil {
		return nil,tx.Error
	}
	return &users, nil
}

func (r *AuthRepo) FetchUserByEmail(email string) (*db.Users, error){
	var user db.Users
	tx:=r.Db.Raw(`SELECT * from users Where email = ?`,email).Scan(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, errors.New(`no user found with this email id`)
		}
	}
	return &user, nil
}