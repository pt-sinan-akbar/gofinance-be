package entities

import (
	"time"
	"github.com/google/uuid"
)

type UserLog struct {
	ID 			int64 		`gorm:"type:int;primary_key;auto_increment" json:"id"`
	UserId 		*uuid.UUID 	`gorm:"type:uuid" json:"user_id,omitempty"`
	Username 	*string 	`gorm:"type:varchar(255)" json:"username,omitempty"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp" json:"created_at,omitempty"`

	
	User 		*User 		`json:"user,omitempty"`
}