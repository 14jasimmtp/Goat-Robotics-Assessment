package interfaceUsecase

import "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"

type ProjectUsecase interface {
	CreateProject(project models.Project, userID int) error
	ListProjects(userID int) ([]models.Project, error)
	AddTask(task models.Task, userID int) error
	DeleteTask(taskID string, userID int) error
	CreateTimeEntry(timeEntry *models.TimeEntry) error
	UpdateTimeEntry(timeEntry *models.TimeEntry, userID int) error
	DeleteTimeEntry(id string, userID int) error
	GetTimeEntryByID(id uint) (*models.TimeEntry, error)
	AddTimeEntry(timeEntry *models.TimeEntry, userID int) error
}
