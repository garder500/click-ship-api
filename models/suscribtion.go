package models

import (
	"gorm.io/gorm"
)

type Suscribtion struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func (s *Suscribtion) TableName() string {
	return "suscribtions.types"
}

type SuscribtionHistory struct {
	gorm.Model
	UserID        uint        `json:"user_id"`
	User          User        `gorm:"foreignKey:UserID"`
	SuscribtionID uint        `json:"suscribtion_id"`
	Suscribtion   Suscribtion `gorm:"foreignKey:SuscribtionID"`
	CardUserID    uint        `json:"card_user_id"`
	CardUser      UserCard    `gorm:"foreignKey:CardUserID"`
}

func (s *SuscribtionHistory) TableName() string {
	return "suscribtions.history"
}
