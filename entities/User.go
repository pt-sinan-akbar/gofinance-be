package entities

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name 		*string 	`gorm:"type:varchar(255)" json:"name,omitempty"`
	Avatar 		*string		`gorm:"type:varchar(255)" json:"avatar,omitempty"`
	Email 		*string 	`gorm:"type:varchar(255)" json:"email,omitempty"`
	Username 	*string 	`gorm:"type:varchar(255)" json:"username,omitempty"`
	Password 	*string 	`gorm:"type:varchar(255)" json:"password,omitempty"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp" json:"updated_at,omitempty"`
	DeletedAt 	*time.Time 	`gorm:"type:timestamp" json:"deleted_at,omitempty"`
}