package entities

import (
	"time"
	"github.com/google/uuid"
)

type Transaction struct {
	ID 				int64 				`gorm:"type:int;primary_key;auto_increment" json:"id"`
	UserId 			uuid.UUID 			`gorm:"type:uuid" json:"user_id,omitempty"`
	Type 			*int32 				`gorm:"type:int" json:"type,omitempty"`
	Amount  		*float64			`gorm:"type:decimal(10,2)" json:"amount,omitempty"`
	Transdate 		*time.Time			`gorm:"type:timestamp" json:"transdate,omitempty"`
	CategoryId  	int32				`gorm:"type:int" json:"category_id,omitempty"`
	Notes 			*string				`gorm:"type:varchar(255)" json:"notes,omitempty"`
	SubCategoryId 	int32				`gorm:"type:int" json:"sub_category_id,omitempty"`
	CreatedAt 		time.Time 			`gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt 		time.Time 			`gorm:"type:timestamp" json:"updated_at,omitempty"`
	DeletedAt 		*time.Time 			`gorm:"type:timestamp" json:"deleted_at,omitempty"`


	User        	User        		`gorm:"foreignkey:UserId" json:"user,omitempty"`
	Category    	Category    		`gorm:"foreignkey:CategoryId" json:"category,omitempty"`
	SubCategory 	SubCategory 		`gorm:"foreignkey:SubCategoryId" json:"sub_category,omitempty"`
}