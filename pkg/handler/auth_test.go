package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/handler"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	mock_interfaceUsecase "github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockAuthUsecase(ctrl)
	handler := &handler.AuthHandler{Usecase: mockUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("successful registration", func(t *testing.T) {
		body := models.Register{
			Email:    "test@example.com",
			Password: "password",
		}

		user := models.RegisterRes{
			Status: "success",
			User: db.Users{Email: "test@example.com"},
			Error: "",
		}

		mockUsecase.EXPECT().Register(body).Return(&user, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Register(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseBody models.RegisterRes
		err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, user.User.Email, responseBody.User.Email)
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_interfaceUsecase.NewMockAuthUsecase(ctrl)
	handler := &handler.AuthHandler{Usecase: mockUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("successful login", func(t *testing.T) {
		body := models.Login{
			Email:    "test@example.com",
			Password: "password",
		}

		user := models.LoginRes{
			Status: "success",
			User: db.Users{Email: "test@example.com"},
		}

		mockUsecase.EXPECT().Login(body).Return(&user, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)

		jsonBody, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Login(c)

		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseBody models.LoginRes
		err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, user.User.Email, responseBody.User.Email)
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