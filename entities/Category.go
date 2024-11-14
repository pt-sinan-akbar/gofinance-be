package entities

import (
	"time"
)

type Category struct {
	ID 			int64 		`gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name 		*string 	`gorm:"type:varchar(255)" json:"name,omitempty"`
	Logo 		*string		`gorm:"type:varchar(255)" json:"avatar,omitempty"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp" json:"updated_at,omitempty"`
	DeletedAt 	*time.Time 	`gorm:"type:timestamp" json:"deleted_at,omitempty"`
}