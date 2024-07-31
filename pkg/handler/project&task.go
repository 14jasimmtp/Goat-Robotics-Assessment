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

// func (h *ProjectHandler) AddTask(c *gin.Context){
// 	var body models.Task

// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err:=h.us.AddTask(body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return 
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "task created successfully"})
// }

// // func (h *ProjectHandler) MarkAsDone(c *gin.Context){

// // }

// func (h *ProjectHandler) DeleteTask(c *gin.Context){
// 	c.Param("id")
// }

// func (h *ProjectHandler) UpdateProject(c *gin.Context){
// 	id:=c.Param("id")
// 	var body models.Project

// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err:=h.us.UpdateProject(body,id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return 
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "project updated successfully"})
// }

// func (h *ProjectHandler) CreateTimeEntry(c *gin.Context){

// }

