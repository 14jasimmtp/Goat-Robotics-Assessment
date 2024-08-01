package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	mockUsecase := new(mocks.AuthUsecase)
	handler := &handler.AuthHandler{us: mockUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("successful registration", func(t *testing.T) {
		body := models.Register{
			Email:    "test@example.com",
			Password: "password",
		}

		user := models.RegisterRes{
			ID:    1,
			Email: body.Email,
		}

		mockUsecase.On("Register", body).Return(user, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Register(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseBody models.User
		err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, responseBody.Email)

		mockUsecase.AssertExpectations(t)
	})

	t.Run("registration with invalid body", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(`invalid json`)))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Register(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestLogin(t *testing.T) {
	mockUsecase := new(mocks.AuthUsecase)
	handler := &handlers.AuthHandler{us: mockUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("successful login", func(t *testing.T) {
		body := models.Login{
			Email:    "test@example.com",
			Password: "password",
		}

		user := models.User{
			ID:    1,
			Email: body.Email,
		}

		mockUsecase.On("Login", body).Return(user, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Login(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseBody models.RegisterRes
		err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, responseBody.Email)

		mockUsecase.AssertExpectations(t)
	})

	t.Run("login with invalid body", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(`invalid json`)))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Login(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
