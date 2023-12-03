package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/frianlh/pokedex-api/libs/constants"
	"github.com/frianlh/pokedex-api/model"
	"gorm.io/gorm"
	"net/http"
)

// MonsterRepositoryInterface is
type MonsterRepositoryInterface interface {
	CreateMonster(tx *gorm.DB, ctx context.Context, req model.Monster) (monsterId string, err error)
	GetMonsterById(ctx context.Context, reqId string, params map[string]interface{}) (res model.Monster, err error)
	GetListMonster(ctx context.Context, queryReq model.MonsterQueryReq, params map[string]interface{}) (res []model.Monster, err error)
	UpdateMonster(tx *gorm.DB, ctx context.Context, req map[string]interface{}) (err error)
	SoftDeleterMonster(tx *gorm.DB, ctx context.Context, req map[string]interface{}) (err error)
	CreateMappingMonsterAndType(tx *gorm.DB, ctx context.Context, req []model.MappingMonsterAndTypes) (err error)
	DeleteMappingMonsterAndType(tx *gorm.DB, ctx context.Context, reqId string) (err error)
	GetLastMonsterCode() (res uint16, err error)
	Transaction() (tx *gorm.DB, resCode int, err error)
}

type monsterRepository struct {
	dbConn *gorm.DB
}

func NewMonsterRepository(db *gorm.DB) MonsterRepositoryInterface {
	return &monsterRepository{
		dbConn: db,
	}
}

// CreateMonster is repository to create monster
func (rMonster *monsterRepository) CreateMonster(tx *gorm.DB, ctx context.Context, req model.Monster) (monsterId string, err error) {
	// transaction
	conn := rMonster.dbConn
	if tx != nil {
		conn = tx
	}

	// create monster
	err = conn.WithContext(ctx).Table(constants.MonsterTable).Create(&req).Error
	if err != nil {
		return "", err
	}

	return req.ID, nil
}

// GetMonsterById is repository to get monster by id
func (rMonster *monsterRepository) GetMonsterById(ctx context.Context, reqId string, params map[string]interface{}) (res model.Monster, err error) {
	query := rMonster.dbConn.WithContext(ctx).Table(constants.MonsterTable)

	// query params
	if params["preloadParams"] != nil {
		for index, _ := range params["preloadParams"].(map[string]interface{}) {
			switch index {
			case "MonsterCategory":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name`)
				})
			case "MonsterTypes":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name`)
				})
			}
		}
	}
	if params["selectParams"] != nil {
		query = query.Select(params["selectParams"])
	}

	// get monster by id
	err = query.Where(`id = ?`, reqId).First(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetListMonster is repository to get list monster
func (rMonster *monsterRepository) GetListMonster(ctx context.Context, queryReq model.MonsterQueryReq, params map[string]interface{}) (res []model.Monster, err error) {
	query := rMonster.dbConn.WithContext(ctx).Table(constants.MonsterTable)

	// query params
	query = query.Order(fmt.Sprintf("monsters.%s %s", queryReq.SortBy, queryReq.OrderBy))
	if params["whereParams"] != nil {
		if params["whereParams"].(map[string]interface{})["default"] != nil {
			for index, value := range params["whereParams"].(map[string]interface{})["default"].(map[string]interface{}) {
				query = query.Where(index, value)
			}
		}
		if params["whereParams"].(map[string]interface{})["in"] != nil {
			for index, value := range params["whereParams"].(map[string]interface{})["in"].(map[string]interface{}) {
				query = query.Where(index, value)
			}
		}
	}
	if params["preloadParams"] != nil {
		for index, _ := range params["preloadParams"].(map[string]interface{}) {
			switch index {
			case "MonsterCategory":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name`)
				})
			case "MonsterTypes":
				query = query.Preload(index, func(query *gorm.DB) *gorm.DB {
					return query.Select(`id, name`)
				})
			}
		}
	}
	if params["joinParams"] != nil {
		for index, value := range params["joinParams"].(map[string]interface{}) {
			if value.(bool) {
				query = query.Joins(index)
			}
		}
	}
	if params["selectParams"] != nil {
		query = query.Select(params["selectParams"])
	}

	// get list monster
	err = query.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateMonster is repository to update monster
func (rMonster *monsterRepository) UpdateMonster(tx *gorm.DB, ctx context.Context, req map[string]interface{}) (err error) {
	// transaction
	conn := rMonster.dbConn
	if tx != nil {
		conn = tx
	}

	query := conn.WithContext(ctx).Table(constants.MonsterTable).Model(&model.Monster{})

	// query params
	if req["whereParams"] != nil {
		if req["whereParams"].(map[string]interface{})["default"] != nil {
			for index, value := range req["whereParams"].(map[string]interface{})["default"].(map[string]interface{}) {
				query = query.Where(index, value)
			}
		}
	}

	// update monster
	err = query.Updates(req["value"]).Error
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleterMonster is repository to soft delete monster
func (rMonster *monsterRepository) SoftDeleterMonster(tx *gorm.DB, ctx context.Context, req map[string]interface{}) (err error) {
	// transaction
	conn := rMonster.dbConn
	if tx != nil {
		conn = tx
	}

	query := conn.WithContext(ctx).Table(constants.MonsterTable).Model(&model.Monster{})

	// query params
	if req["whereParams"] != nil {
		if req["whereParams"].(map[string]interface{})["default"] != nil {
			for index, value := range req["whereParams"].(map[string]interface{})["default"].(map[string]interface{}) {
				query = query.Where(index, value)
			}
		}
	}

	// soft delete monster
	err = query.Delete(&model.Monster{}).Error
	if err != nil {
		return err
	}

	return nil
}

// CreateMappingMonsterAndType is repository to create mapping monster and monster type
func (rMonster *monsterRepository) CreateMappingMonsterAndType(tx *gorm.DB, ctx context.Context, req []model.MappingMonsterAndTypes) (err error) {
	// transaction
	conn := rMonster.dbConn
	if tx != nil {
		conn = tx
	}

	// create mapping monster and monster type
	err = conn.WithContext(ctx).Table(constants.MappingMonsterAndTypes).Create(&req).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteMappingMonsterAndType is repository to delete mapping monster and monster type based on monster id
func (rMonster *monsterRepository) DeleteMappingMonsterAndType(tx *gorm.DB, ctx context.Context, reqId string) (err error) {
	// transaction
	conn := rMonster.dbConn
	if tx != nil {
		conn = tx
	}

	// delete mapping monster and monster type
	err = conn.WithContext(ctx).Table(constants.MappingMonsterAndTypes).
		Where(`monster_id = ?`, reqId).
		Delete(&model.MappingMonsterAndTypes{}).Error
	if err != nil {
		return err
	}

	return nil
}

// GetLastMonsterCode is repository to get last monster code
func (rMonster *monsterRepository) GetLastMonsterCode() (res uint16, err error) {
	// get last monster code
	err = rMonster.dbConn.Table(constants.MonsterTable).
		Select(`monster_code`).
		Model(&model.Monster{}).
		Order(`created_at DESC`).
		First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return res, err
	}

	return res, nil
}

// Transaction is repository to create transactional database
func (rMonster *monsterRepository) Transaction() (tx *gorm.DB, resCode int, err error) {
	return rMonster.dbConn, http.StatusInternalServerError, nil
}
