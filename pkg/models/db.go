package models

import (
	"time"
)

type TimeEntry struct {
	ID        uint `gorm:"primaryKey"`
	TaskID    uint
	UserID    uint `gorm:"not null"`
	StartTime time.Time
	EndTime   time.Time
}
