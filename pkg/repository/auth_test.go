package repository_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/models"
	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func TestSaveUser(t *testing.T) {
	t.Run("successful save user", func(t *testing.T) {
		// Setup sqlmock
		dbMock, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer dbMock.Close()

		gormDB, _, err := setupMockDB()
		assert.NoError(t, err)

		// Setup the repository
		authRepo := repository.NewAuthRepo(gormDB)

		// Set up the expected insert
		mock.ExpectExec(`INSERT INTO "users" \("username","email","password"\) VALUES \(\$1,\$2,\$3\)`).
			WithArgs("testuser", "test@example.com", "hashedpassword").
			WillReturnResult(sqlmock.NewResult(1, 1))

		user := models.Register{
			Username: "testuser",
			Email:    "test@example.com",
		}

		savedUser, err := authRepo.SaveUser(user, "hashedpassword")

		assert.NoError(t, err)
		assert.NotNil(t, savedUser)
		assert.Equal(t, "testuser", savedUser.Username)
		assert.Equal(t, "test@example.com", savedUser.Email)
		assert.Equal(t, "hashedpassword", savedUser.Password)

		// Ensure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
func TestFetchUserByEmail(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	authRepo := repository.NewAuthRepo(gormDB)

	t.Run("successful fetch user by email", func(t *testing.T) {
		email := "test@example.com"
		user := db.Users{
			Username: "testuser",
			Email:    email,
			Password: "hashedpassword",
		}

		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(user.ID, user.Username, user.Email, user.Password)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * from users Where email = ?`)).
			WithArgs(email).
			WillReturnRows(rows)

		fetchedUser, err := authRepo.FetchUserByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedUser)
		assert.Equal(t, user.ID, fetchedUser.ID)
		assert.Equal(t, user.Username, fetchedUser.Username)
		assert.Equal(t, user.Email, fetchedUser.Email)
		assert.Equal(t, user.Password, fetchedUser.Password)
	})

	t.Run("fetch user by email with no record found", func(t *testing.T) {
		email := "test@example.com"

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * from users Where email = ?`)).
			WithArgs(email).
			WillReturnError(gorm.ErrRecordNotFound)

		fetchedUser, err := authRepo.FetchUserByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, fetchedUser)
		assert.Equal(t, "no user found with this email id", err.Error())
	})

	t.Run("fetch user by email with query error", func(t *testing.T) {
		email := "test@example.com"

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * from users Where email = ?`)).
			WithArgs(email).
			WillReturnError(errors.New("query error"))

		fetchedUser, err := authRepo.FetchUserByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, fetchedUser)
	})
}
