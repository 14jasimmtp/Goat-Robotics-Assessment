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
	ID        uint `gorm:"primaryKey"`
	TaskID    uint
	UserID    uint  `gorm:"not null"`
	User      Users `gorm:"foreignkey:UserID"`
	StartTime time.Time
	EndTime   time.Time
}

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})

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

	db, err = gorm.Open(postgres.Open(viper.GetString("DB_URL")+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
	 db.AutoMigrate(&TimeEntry{})
	return db
}
