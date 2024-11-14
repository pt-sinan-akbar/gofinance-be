package entities

import (
	"time"
)

type SubCategory struct {
	ID 			int64 		`gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name 		*string 	`gorm:"type:varchar(255)" json:"name,omitempty"`
	Logo 		*string 	`gorm:"type:varchar(255)" json:"avatar,omitempty"`
	CategoryId 	*int32 		`gorm:"type:int" json:"category_id,omitempty"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp" json:"updated_at,omitempty"`
	DeletedAt 	*time.Time 	`gorm:"type:timestamp" json:"deleted_at,omitempty"`


	Category 	*Category	`json:"category,omitempty"`
}