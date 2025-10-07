package database

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
