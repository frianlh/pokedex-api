package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	ID        string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name      string          `json:"name"`
	Action    string          `json:"action"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (Permission) TableName() string {
	return constants.PermissionTable
}
