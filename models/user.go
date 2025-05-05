package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Role      string    `gorm:"type:varchar(10)" json:"role"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(50);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     []Task    `gorm:"constraint:OnDelete:CASCADE;" json:"tasks,omitempty"`
}
