package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	interfaceUsecase "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase/inteface"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	us interfaceUsecase.ProjectUsecase
}

func NewProjectHandler(us interfaceUsecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{us: us}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var body models.Project
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		fmt.Println("user", userID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found"})
		return
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter details correctly"})
		return
	}

	err := h.us.CreateProject(body, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "project created successfully"})
}

func (h *ProjectHandler) ListProjects(c *gin.Context) {
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}
	projects, err := h.us.ListProjects(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) AddTask(c *gin.Context) {
	var body models.Task
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.us.AddTask(body, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "task added successfully"})
}

func (h *ProjectHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}

	err := h.us.DeleteTask(taskID, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "task deleted successfully"})
}

func (h *ProjectHandler) AddTimeEntry(c *gin.Context) {
	var body models.TimeEntry
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.us.AddTimeEntry(&body, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "time entry added successfully"})
}

func (h *ProjectHandler) UpdateTimeEntry(c *gin.Context) {
	var timeEntry models.TimeEntry
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&timeEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ID, _ := strconv.Atoi(id)
	timeEntry.ID = uint(ID)
	if err := h.us.UpdateTimeEntry(&timeEntry, int(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, timeEntry)
}

func (h *ProjectHandler) DeleteTimeEntry(c *gin.Context) {
	id := c.Param("id")
	userID, exist := c.MustGet("User_id").(uint)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not authorised"})
		return
	}
	if err := h.us.DeleteTimeEntry(id, int(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "time entry deleted"})
}
