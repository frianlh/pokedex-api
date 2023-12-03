package repository

import (
	"context"
	"github.com/frianlh/pokedex-api/libs/constants"
	"github.com/frianlh/pokedex-api/model"
	"gorm.io/gorm"
)

// MTypeRepositoryInterface is
type MTypeRepositoryInterface interface {
	GetAllMonsterType(ctx context.Context, selectParams []string) (res []model.MonsterType, err error)
}

type mTypeRepository struct {
	dbConn *gorm.DB
}

func NewMTypeRepository(db *gorm.DB) MTypeRepositoryInterface {
	return &mTypeRepository{
		dbConn: db,
	}
}

// GetAllMonsterType is repository to get all monster type based on select params
func (rMType *mTypeRepository) GetAllMonsterType(ctx context.Context, selectParams []string) (res []model.MonsterType, err error) {
	query := rMType.dbConn.WithContext(ctx).Table(constants.MonsterTypeTable)

	// query params
	if selectParams != nil {
		query = query.Select(selectParams)
	}

	// get all monster type
	err = query.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
