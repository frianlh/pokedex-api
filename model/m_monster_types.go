package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type MonsterType struct {
	ID        string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (MonsterType) TableName() string {
	return constants.MonsterTypeTable
}
