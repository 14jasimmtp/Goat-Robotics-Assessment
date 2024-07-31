package usecase

import (
	"errors"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository"
)

type ProjectUsecase struct {
	Repo repository.ProjectRepo 
}

func NewProjectUsecase(repo repository.ProjectRepo) ProjectUsecase {
	return ProjectUsecase{Repo: repo}
}

func (u *ProjectUsecase) CreateProject(project models.Project) (error){
	if u.Repo.CheckProjectExist(project) {
		return errors.New(`project already exist with this name`)
	}

	err := u.Repo.CreateProject(project)
	if err != nil {
		return err
	}
	return nil
}