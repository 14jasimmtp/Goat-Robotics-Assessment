package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	mock_interfaceUsecase "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful project creation", func(t *testing.T) {
		body := models.Project{Name: "New Project"}
		userID := uint(1)

		mockUsecase.EXPECT().CreateProject(body, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/project", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)

		handler.CreateProject(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "project created successfully")
	})

	t.Run("project creation with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request, _ = http.NewRequest(http.MethodPost, "/project", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.CreateProject(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "no user found")
	})
}

func TestListProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful list projects", func(t *testing.T) {
		userID := uint(1)
		projects := []models.Project{{Name: "Project 1"}, {Name: "Project 2"}}

		mockUsecase.EXPECT().ListProjects(int(userID)).Return(projects, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Set("User_id", userID)

		handler.ListProjects(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseProjects []models.Project
		err := json.Unmarshal(recorder.Body.Bytes(), &responseProjects)
		assert.NoError(t, err)
		assert.Equal(t, projects, responseProjects)
	})

	t.Run("list projects with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		handler.ListProjects(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful add task", func(t *testing.T) {
		body := models.Task{ProjectID: 1, Name: "New Task"}
		userID := uint(1)

		mockUsecase.EXPECT().AddTask(body, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)

		handler.AddTask(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "task added successfully")
	})

	t.Run("add task with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request, _ = http.NewRequest(http.MethodPost, "/task", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.AddTask(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful delete task", func(t *testing.T) {
		taskID := "1"
		userID := uint(1)

		mockUsecase.EXPECT().DeleteTask(taskID, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request, _ = http.NewRequest(http.MethodDelete, "/task/"+taskID, nil)
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)
		c.AddParam("id", taskID)

		handler.DeleteTask(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "task deleted successfully")
	})

	t.Run("delete task with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request, _ = http.NewRequest(http.MethodDelete, "/task/1", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.DeleteTask(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}

func TestAddTimeEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful add time entry", func(t *testing.T) {
		body := models.TimeEntry{TaskID: 1}
		userID := uint(1)

		mockUsecase.EXPECT().AddTimeEntry(&body, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/time-entry", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)

		handler.AddTimeEntry(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "time entry added successfully")
	})

	t.Run("add time entry with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request, _ = http.NewRequest(http.MethodPost, "/time-entry", nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.AddTimeEntry(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}

func TestUpdateTimeEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful update time entry", func(t *testing.T) {
		timeEntry := models.TimeEntry{TaskID: 1, ID: 1}
		userID := uint(1)
		id := "1"

		mockUsecase.EXPECT().UpdateTimeEntry(&timeEntry, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(timeEntry)
		c.Request, _ = http.NewRequest(http.MethodPut, "/time-entry/"+id, bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)
		c.AddParam("id", id)

		handler.UpdateTimeEntry(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseTimeEntry models.TimeEntry
		err := json.Unmarshal(recorder.Body.Bytes(), &responseTimeEntry)
		assert.NoError(t, err)
		assert.Equal(t, timeEntry, responseTimeEntry)
	})

	t.Run("update time entry with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		id := "1"

		c.Request, _ = http.NewRequest(http.MethodPut, "/time-entry/"+id, nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.UpdateTimeEntry(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}

func TestDeleteTimeEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockProjectUsecase(ctrl)
	handler := handler.NewProjectHandler(mockUsecase)

	gin.SetMode(gin.TestMode)

	t.Run("successful delete time entry", func(t *testing.T) {
		id := "1"
		userID := uint(1)

		mockUsecase.EXPECT().DeleteTimeEntry(id, int(userID)).Return(nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		c.Request, _ = http.NewRequest(http.MethodDelete, "/time-entry/"+id, nil)
		c.Request.Header.Set("Content-Type", "application/json")
		c.Set("User_id", userID)
		c.AddParam("id", id)

		handler.DeleteTimeEntry(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "time entry deleted")
	})

	t.Run("delete time entry with missing user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		id := "1"

		c.Request, _ = http.NewRequest(http.MethodDelete, "/time-entry/"+id, nil)
		c.Request.Header.Set("Content-Type", "application/json")

		handler.DeleteTimeEntry(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "not authorised")
	})
}
