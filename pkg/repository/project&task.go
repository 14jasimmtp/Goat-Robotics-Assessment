package repository

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	Db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) AuthRepo {
	return AuthRepo{Db: db}
}
func (u *ProjectRepo) CheckProjectExist(project models.Project) bool {
	var val bool
	u.Db.Raw(`SELECT EXISTS(SELECT * From projects where name = ? AND client = ?)`, project.Name, project.Client).Scan(&val)
	return val
}

func (u *ProjectRepo) CreateProject(project models.Project) error {
	tx := u.Db.Save(&db.Project{Name: project.Name, Client: project.Client, Description: project.Description})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) UpdateProject(project models.Project) error {
	tx := u.Db.Raw(`UPDATE projects SET name = ?,description = ?, client = ?`)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) AddTask(task models.Task) error {
	tx := u.Db.Save(&db.Task{Name: task.Name,ProjectID: task.ProjectID})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) DeleteTask() {
	
}

// func (u *ProjectRepo)
