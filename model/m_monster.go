package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"time"
)

type Monster struct {
	ID                string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	MonsterCode       uint16          `json:"monster_code"`
	Name              string          `json:"name"`
	MonsterCategoryId string          `json:"monster_category_id"`
	MonsterCategory   MonsterCategory `json:"monster_category" gorm:"foreignKey:MonsterCategoryId;references:ID"`
	MonsterTypes      []MonsterType   `json:"monster_types" gorm:"many2many:mapping_monster_and_types;save_association:false"`
	Description       string          `json:"description"`
	Length            float32         `json:"length"`
	Weight            uint16          `json:"weight"`
	HP                uint16          `json:"hp"`
	Attack            uint16          `json:"attack"`
	Defends           uint16          `json:"defends"`
	Speed             uint16          `json:"speed"`
	IsCaught          bool            `json:"is_caught"`
	ImageName         string          `json:"image_name"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at"`
}

func (Monster) TableName() string {
	return constants.MonsterTable
}

type MonsterQueryReq struct {
	SortBy        string   `json:"sort_by"`
	OrderBy       string   `json:"order_by"`
	Name          string   `json:"name"`
	MonsterTypeId []string `json:"monster_type_id"`
	IsCaught      string   `json:"is_caught"`
}

type CreateMonsterReq struct {
	Name              string   `json:"name" form:"name" validate:"required"`
	MonsterCategoryId string   `json:"monster_category_id" form:"monster_category_id" validate:"required"`
	MonsterTypes      []string `json:"monster_types" form:"monster_types"`
	Description       string   `json:"description" form:"description" validate:"required"`
	Length            float32  `json:"length" form:"length" validate:"required"`
	Weight            uint16   `json:"weight" form:"weight" validate:"required"`
	HP                uint16   `json:"hp" form:"hp" validate:"required"`
	Attack            uint16   `json:"attack" form:"attack" validate:"required"`
	Defends           uint16   `json:"defends" form:"defends" validate:"required"`
	Speed             uint16   `json:"speed" form:"speed" validate:"required"`
	ImageName         string   `json:"image_name"`
}

type GetListMonsterRes struct {
	ID              string          `json:"id"`
	MonsterCode     uint16          `json:"monster_code"`
	Name            string          `json:"name"`
	MonsterCategory MonsterCategory `json:"monster_category"`
	MonsterTypes    []MonsterType   `json:"monster_types"`
	IsCaught        bool            `json:"is_caught"`
	ImageName       string          `json:"image_name"`
	ImageURL        string          `json:"image_url"`
}

type GetDetailMonsterRes struct {
	ID              string          `json:"id" gorm:"unique;default:gen_random_uuid()"`
	MonsterCode     uint16          `json:"monster_code"`
	Name            string          `json:"name"`
	MonsterCategory MonsterCategory `json:"monster_category" gorm:"foreignKey:MonsterCategoryId;references:ID"`
	MonsterTypes    []MonsterType   `json:"monster_types" gorm:"many2many:mapping_monster_and_types;save_association:false"`
	Description     string          `json:"description"`
	Length          float32         `json:"length"`
	Weight          uint16          `json:"weight"`
	HP              uint16          `json:"hp"`
	Attack          uint16          `json:"attack"`
	Defends         uint16          `json:"defends"`
	Speed           uint16          `json:"speed"`
	IsCaught        bool            `json:"is_caught"`
	ImageName       string          `json:"image_name"`
	ImageURL        string          `json:"image_url"`
}

type UpdateMonsterReq struct {
	Name              string   `json:"name" form:"name"`
	MonsterCategoryId string   `json:"monster_category_id" form:"monster_category_id"`
	MonsterTypes      []string `json:"monster_types" form:"monster_types"`
	Description       string   `json:"description" form:"description"`
	Length            float32  `json:"length" form:"length"`
	Weight            uint16   `json:"weight" form:"weight"`
	HP                uint16   `json:"hp" form:"hp"`
	Attack            uint16   `json:"attack" form:"attack"`
	Defends           uint16   `json:"defends" form:"defends"`
	Speed             uint16   `json:"speed" form:"speed"`
	IsCaught          bool     `json:"is_caught" form:"is_caught"`
	ImageName         string   `json:"image_name"`
}

type UpdateMonsterCapturedReq struct {
	IsCaught bool `json:"is_caught" form:"is_caught"`
}
