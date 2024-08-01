package interfaceRepo

import "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"

type ProjectRepository interface {
	CheckProjectExist(project models.Project) bool
    CreateProject(project models.Project, userID int) error
    UpdateProject(project models.Project) error
    CheckProjectExistByID(id uint, userID uint) (bool, error)
    CheckTaskExistInProject(projectID uint, task string) (bool, error)
    AddTask(task models.Task) error
    DeleteTask(taskID string, userid int) error
    CheckTaskExistByID(taskID string, userID int) (bool, error)
    CreateTimeEntry(timeEntry *models.TimeEntry) error
    UpdateTimeEntry(timeEntry *models.TimeEntry, UserID int) error
    DeleteTimeEntry(id string, userID int) error
    GetByIDTimeEntry(id uint) (*models.TimeEntry, error)
    GetAll() ([]models.TimeEntry, error)
    ListProjects(userID int) ([]models.Project, error)
}
