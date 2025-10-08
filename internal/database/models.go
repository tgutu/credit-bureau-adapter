package database

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreditBureau struct {
	BaseModel
	Name string `gorm:"uniqueIndex"`
}

func (b *CreditBureau) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	b.ID = id
	return nil
}
