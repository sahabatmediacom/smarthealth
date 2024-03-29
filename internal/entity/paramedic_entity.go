package entity

import (
	"time"

	"gorm.io/gorm"
)

type Paramedic struct {
	gorm.Model
	Name        string
	Hospitals   []Hospital `gorm:"many2many:paramedic_hospitals;"`
	UserID      int
	IDSatusehat string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
