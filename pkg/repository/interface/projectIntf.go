package interfaceRepo

import "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"

type ProjectRepository interface {
	CheckProjectExist(project models.Project) bool
	CreateProject(project models.Project, userID int) error
	UpdateProject(project models.Project) error
	CheckProjectExistByID(id uint) (bool, error)
	CheckTaskExistInProject(projectID uint, task string) (bool, error)
	AddTask(task models.Task) error
	DeleteTask(taskID uint) error
	CheckTaskExistByID(taskID uint, userID int) (bool, error)
	AddTimeEntry(timeEntry models.TimeEntry) error
}
