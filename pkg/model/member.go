package model

import (
	"time"
)

type Member struct {
	ID        string     `json:"id" db:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string     `json:"name" db:"name" gorm:"type:text;not null"`
	Username  string     `json:"username" db:"username" gorm:"type:text;not null"`
	Password  string     `json:"password" db:"password" gorm:"type:text;not null"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at" gorm:"index"`
}

func (Member) TableName() string {
	return "member"
}
