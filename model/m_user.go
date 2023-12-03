package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name              string          `json:"name"`
	Email             string          `json:"email"`
	EncryptedPassword string          `json:"-"`
	RoleId            string          `json:"role_id"`
	Role              *Role           `json:"role" gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at"`
}

func (User) TableName() string {
	return constants.UserTable
}
