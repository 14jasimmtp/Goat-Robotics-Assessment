package db

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string
	Username string
	Password string
}

type Project struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Client      string
	Description string
	CreatedBy   uint  `gorm:"not null"`
	User        Users `gorm:"foreignkey:CreatedBy"`
}

type Task struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	ProjectID   uint    `gorm:"not null"`
	Project     Project `gorm:"foreignkey:ProjectID"`
}

type TimeEntry struct {
	gorm.Model
	TaskID    uint      `gorm:"not null"`
	Task      Task      `gorm:"foreignkey:TaskID"`
	UserID    uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
	Duration  time.Duration
}

func ConnectToDB() *gorm.DB {
	fmt.Println(viper.GetString("DB_URL"),"jelllo")
	db, err := gorm.Open(postgres.Open("postgres://postgres:jasi123@localhost:5432"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "timetracker"

	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		err = db.Exec("CREATE DATABASE timetracker").Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	db, err = gorm.Open(postgres.Open("postgres://postgres:jasi123@localhost:5432"+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&TimeEntry{})
	return db
}
