package repository

import (
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	interfaceRepo "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository/interface"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	Db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) interfaceRepo.ProjectRepository {
	return &ProjectRepo{Db: db}
}

func (u *ProjectRepo) CheckProjectExist(project models.Project) bool {
	var val bool
	u.Db.Raw(`SELECT EXISTS(SELECT * FROM projects WHERE name = ? AND client = ?)`, project.Name, project.Client).Scan(&val)
	return val
}

func (u *ProjectRepo) CreateProject(project models.Project, userID int) error {
	pro:=db.Project{Name: project.Name,Description: project.Description,Client: project.Client,CreatedBy: uint(userID)}
	tx := u.Db.Create(&pro)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) UpdateProject(project models.Project) error {
	tx := u.Db.Model(&project).Updates(models.Project{Name: project.Name, Description: project.Description, Client: project.Client})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) CheckProjectExistByID(id uint) (bool, error) {
	var count int64
	err := u.Db.Model(&models.Project{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *ProjectRepo) CheckTaskExistInProject(projectID uint, task string) (bool, error) {
	var count int64
	err := u.Db.Model(&models.Task{}).Where("name = ? AND project_id = ?", task, projectID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *ProjectRepo) AddTask(task models.Task) error {
	tx := u.Db.Create(&task)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) DeleteTask(taskID uint) error {
	tx := u.Db.Delete(&models.Task{}, taskID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *ProjectRepo) CheckTaskExistByID(taskID uint, userID int)(bool,error) {

}

func (r *ProjectRepo) CreateTimeEntry(timeEntry *models.TimeEntry) error {
    return r.Db.Create(timeEntry).Error
}

func (r *ProjectRepo) UpdateTimeEntry(timeEntry *models.TimeEntry) error {
    return r.Db.Save(timeEntry).Error
}

func (r *ProjectRepo) DeleteTimeEntry(id uint) error {
    return r.Db.Delete(&models.TimeEntry{}, id).Error
}

func (r *ProjectRepo) GetByIDTimeEntry(id uint) (*models.TimeEntry, error) {
    var timeEntry models.TimeEntry
    err := r.Db.First(&timeEntry, id).Error
    return &timeEntry, err
}

func (r *ProjectRepo) GetAll() ([]models.TimeEntry, error) {
    var timeEntries []models.TimeEntry
    err := r.Db.Find(&timeEntries).Error
    return timeEntries, err
}