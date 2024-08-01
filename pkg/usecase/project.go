package usecase

import (
	"errors"
	"strconv"
	"time"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	interfaceRepo "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository/interface"
)

type ProjectUsecase struct {
	Repo interfaceRepo.ProjectRepository
}

func NewProjectUsecase(repo interfaceRepo.ProjectRepository) ProjectUsecase {
	return ProjectUsecase{Repo: repo}
}

func (u *ProjectUsecase) CreateProject(project models.Project, userID int) error {
	if u.Repo.CheckProjectExist(project) {
		return errors.New(`project already exist with this name`)
	}

	err := u.Repo.CreateProject(project, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProjectUsecase) ListProjects(userID int) ([]models.Project, error) {
	return u.Repo.ListProjects(userID)
}

func (u *ProjectUsecase) AddTask(task models.Task, UserID int) error {
	exist, err := u.Repo.CheckProjectExistByID(task.ProjectID, uint(UserID))
	if err != nil {
		return errors.New(`something went wrong 1`)
	}

	if !exist {
		return errors.New(`project doesn't exist with this id`)
	}

	exist, err = u.Repo.CheckTaskExistInProject(task.ProjectID, task.Name)
	if err != nil {
		return errors.New(`something went wrong 2`)
	}
	if exist {
		return errors.New(`task already added in this project`)
	}

	err = u.Repo.AddTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProjectUsecase) DeleteTask(taskID string, userID int) error {
	exist, err := u.Repo.CheckTaskExistByID(taskID, userID)
	if err != nil {
		return errors.New("something went wrong")
	}

	if !exist {
		return errors.New("task does not exist")
	}

	err = u.Repo.DeleteTask(taskID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProjectUsecase) CreateTimeEntry(timeEntry *models.TimeEntry) error {
	timeEntry.StartTime = time.Now()
	return u.Repo.CreateTimeEntry(timeEntry)
}

func (u *ProjectUsecase) UpdateTimeEntry(timeEntry *models.TimeEntry, userID int) error {
	timeEntry.EndTime = time.Now()
	return u.Repo.UpdateTimeEntry(timeEntry, userID)
}

func (u *ProjectUsecase) DeleteTimeEntry(id string, userID int) error {
	return u.Repo.DeleteTimeEntry(id, userID)
}

func (u *ProjectUsecase) GetTimeEntryByID(id uint) (*models.TimeEntry, error) {
	return u.Repo.GetByIDTimeEntry(id)
}

func (u *ProjectUsecase) AddTimeEntry(timeEntry *models.TimeEntry, userID int) error {
	exist, err := u.Repo.CheckTaskExistByID(strconv.Itoa(int(timeEntry.TaskID)), userID)
	if err != nil {
		return errors.New("something went wrong")
	}

	if !exist {
		return errors.New("task does not exist")
	}
	timeEntry.StartTime = time.Now()
	err = u.Repo.CreateTimeEntry(timeEntry)
	if err != nil {
		return err
	}
	return nil
}
