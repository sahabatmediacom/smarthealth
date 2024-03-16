package entity

import (
	"time"

	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name       string
	IP         string
	Paramedics []Paramedic `gorm:"many2many:paramedic_hospitals;"`
	Xusername  string
	Xpassword  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
