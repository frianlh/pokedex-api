package repository

import (
	"context"
	"github.com/frianlh/pokedex-api/libs/constants"
	"github.com/frianlh/pokedex-api/model"
	"gorm.io/gorm"
)

// MCategoryRepositoryInterface is
type MCategoryRepositoryInterface interface {
	GetAllMonsterCategory(ctx context.Context, selectParams []string) (res []model.MonsterCategory, err error)
}

type mCategoryRepository struct {
	dbConn *gorm.DB
}

func NewMonsterCategory(db *gorm.DB) MCategoryRepositoryInterface {
	return &mCategoryRepository{
		dbConn: db,
	}
}

// GetAllMonsterCategory is repository to get all monster category based on select params
func (rMCategory *mCategoryRepository) GetAllMonsterCategory(ctx context.Context, selectParams []string) (res []model.MonsterCategory, err error) {
	query := rMCategory.dbConn.WithContext(ctx).Table(constants.MonsterCategoryTable)

	// query params
	if selectParams != nil {
		query = query.Select(selectParams)
	}

	// get all monster category
	err = query.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
