package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID        int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Role      string    `gorm:"type:varchar(10)" json:"role"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(50)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Tasks     []Task    `gorm:"constraint:OnDelete:CASCADE;" json:"tasks,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u User) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Clauses(clause.Returning{}).Where("user_id = ?", u.ID).Delete(&Task{}).Error
}
