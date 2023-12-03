package model

import (
	"github.com/frianlh/pokedex-api/libs/constants"
)

type MappingMonsterAndTypes struct {
	MonsterId     string `json:"monster_id"`
	MonsterTypeId string `json:"monster_type_id"`
}

func (MappingMonsterAndTypes) TableName() string {
	return constants.MappingMonsterAndTypes
}
