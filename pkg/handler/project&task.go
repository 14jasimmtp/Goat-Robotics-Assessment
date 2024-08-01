package handler

import (
	"net/http"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct{
	us usecase.ProjectUsecase
}

func NewProjectHandler(us usecase.ProjectUsecase) *ProjectHandler{
	return &ProjectHandler{us: us}
}

func (h *ProjectHandler) CreateProject(c *gin.Context){
	var body models.Project

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err:=h.us.CreateProject(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"status": "project created successfully"})

}

func (h *ProjectHandler) ListProjects(c *gin.Context){

}

func (h *ProjectHandler) AddTask(c *gin.Context){
	var body models.Task

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err:=h.us.AddTask(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
}

func (h *ProjectHandler) DeleteTask(c *gin.Context){
	taskID := c.Param("id")
	user_id,exist:=c.Get("User_id")
	if !exist{
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}
	
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}

	err := h.us.DeleteTask(taskID, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"status": "task deleted successfully"})
}


func (h *ProjectHandler) AddTimeEntry(c *gin.Context){
	var body models.TimeEntry

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.us.AddTimeEntry(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"status": "time entry added successfully"})
}

func (h *ProjectHandler) CreateTimeEntry(c *gin.Context) {
    var timeEntry models.TimeEntry
    if err := c.ShouldBindJSON(&timeEntry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.usecase.CreateTimeEntry(&timeEntry); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, timeEntry)
}

func (h *ProjectHandler) UpdateTimeEntry(c *gin.Context) {
    var timeEntry models.TimeEntry
    id := c.Param("id")
    if err := c.ShouldBindJSON(&timeEntry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    timeEntry.ID = id
    if err := h.usecase.UpdateTimeEntry(&timeEntry); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, timeEntry)
}

func (h *ProjectHandler) DeleteTimeEntry(c *gin.Context) {
    id := c.Param("id")
    if err := h.usecase.DeleteTimeEntry(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Time entry deleted"})
}

func (h *ProjectHandler) GetTimeEntries(c *gin.Context) {
    timeEntries, err := h.usecase.GetAllTimeEntries()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, timeEntries)
}