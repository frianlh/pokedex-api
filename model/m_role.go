package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name        string          `json:"name"`
	Permissions []Permission    `json:"permissions" gorm:"many2many:role_permissions;save_association:false"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
}

func (Role) TableName() string {
	return constants.RoleTable
}
