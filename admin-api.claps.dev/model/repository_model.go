package model

import "time"

type Repository struct {
	Id          string    `json:"id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	ProjectId   string    `json:"project_id,omitempty" gorm:"type:varchar(50);not null"`
	Type        string    `json:"type,omitempty" gorm:"varchar(10);not null"`
	Name        string    `json:"name,omitempty" gorm:"type:varchar(50);not null"`
	Slug        string    `json:"slug,omitempty" gorm:"type:varchar(100);not null"`
	Description string    `json:"description,omitempty" gorm:"type:varchar(120);default:null"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"type:timestamp with time zone"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"type:timestamp with time zone"`
}
