package repository

import (
	"context"
	"github.com/frianlh/pokedex-api/libs/constants"
	"github.com/frianlh/pokedex-api/model"
	"gorm.io/gorm"
)

// UserRepositoryInterface is
type UserRepositoryInterface interface {
	GetUserByParams(ctx context.Context, params map[string]interface{}) (res model.User, err error)
}

type userRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		dbConn: db,
	}
}

// GetUserByParams is repository to get user by params
func (rUser *userRepository) GetUserByParams(ctx context.Context, params map[string]interface{}) (res model.User, err error) {
	query := rUser.dbConn.WithContext(ctx).Table(constants.UserTable)

	// query params
	if params["whereParams"] != nil {
		if params["whereParams"].(map[string]interface{})["default"] != nil {
			for index, value := range params["whereParams"].(map[string]interface{})["default"].(map[string]interface{}) {
				query = query.Where(index, value)
			}
		}
	}
	if params["preloadParams"] != nil {
		for index, _ := range params["preloadParams"].(map[string]interface{}) {
			switch index {
			case "Role":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name`)
				})
			case "Role.Permissions":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name, action`)
				})
			}
		}
	}
	if params["selectParams"] != nil {
		query = query.Select(params["selectParams"])
	}

	// get user by params
	err = query.First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
