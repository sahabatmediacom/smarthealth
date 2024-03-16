package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Password    string
	ParamedicID string
	// Paramedic   Paramedic `gorm:"foreignKey:ParamedicID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
