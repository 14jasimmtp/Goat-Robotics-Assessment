package models

type Register struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Client      string `json:"client"`
}

type Task struct {
	Name      string `json:"name"`
	ProjectID uint   `json:"project_id"`
}
