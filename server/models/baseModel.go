package models

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
//    type User struct {
//      models.BaseModel
//    }
type BaseModel struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
