package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type MonsterCategory struct {
	ID        string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (MonsterCategory) TableName() string {
	return constants.MonsterCategoryTable
}
