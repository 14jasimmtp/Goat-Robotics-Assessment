package models

import "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"

type RegisterRes struct{
	Status string	`json:"status"`
	User db.Users	`json:"user,omitempty"`
	Error string `json:"error,omitempty"`
}

type LoginRes struct{
	Status string `json:"status"`
	User db.Users `json:"user,omitempty"`
	Token string `json:"access_token"`
	Error string `json:"error,omitempty"`
}
