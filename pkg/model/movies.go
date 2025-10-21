package model

import (
	"time"
)

type Movie struct {
	ID        string     `json:"id" db:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string     `json:"name" db:"name" gorm:"type:text;not null"`
	Star      int        `json:"star" db:"star" gorm:"type:int"`
	Actor     string     `json:"actor" db:"actor" gorm:"type:text"`
	CreatedBy string     `json:"created_by" db:"created_by" gorm:"type:varchar(36)"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at" gorm:"index"`

	// Relationship
	Creator *Member `json:"creator,omitempty" gorm:"foreignKey:CreatedBy;references:ID"`
}
