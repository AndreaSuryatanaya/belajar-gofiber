package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Todo      string    `json:"todo" gorm:"not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) error {
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}
	return nil
}
