package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/frianlh/pokedex-api/libs/uploader"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/repository"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// MonsterUseCaseInterface is
type MonsterUseCaseInterface interface {
	CreateMonster(ctx context.Context, req model.CreateMonsterReq) (resCode int, resMessage string, err error)
	GetMonsterById(ctx context.Context, reqId string) (res model.GetDetailMonsterRes, resCode int, resMessage string, err error)
	GetListMonster(ctx context.Context, queryReq model.MonsterQueryReq) (res []model.GetListMonsterRes, resCode int, resMessage string, err error)
	UpdateMonster(ctx context.Context, reqId string, req model.UpdateMonsterReq) (resCode int, resMessage string, err error)
	UpdateMonsterCaptured(ctx context.Context, reqId string, req model.UpdateMonsterCapturedReq) (resCode int, resMessage string, err error)
	DeleteMonster(ctx context.Context, reqId string) (resCode int, resMessage string, err error)
}

type monsterUseCase struct {
	ctxTimeout  time.Duration
	baseURL     string
	monsterRepo repository.MonsterRepositoryInterface
}

func NewMonsterUseCase(ctxTimeout time.Duration, baseURL string, monsterRepo repository.MonsterRepositoryInterface) MonsterUseCaseInterface {
	return &monsterUseCase{
		ctxTimeout:  ctxTimeout,
		baseURL:     baseURL,
		monsterRepo: monsterRepo,
	}
}

// CreateMonster is use case to create monster
func (uMonster *monsterUseCase) CreateMonster(ctx context.Context, req model.CreateMonsterReq) (resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	var tx = &gorm.DB{}
	defer func() {
		if rec := recover(); rec != nil {
			// mapping response data
			resCode = http.StatusInternalServerError
			resMessage = "failed create monster"
			err = fmt.Errorf("%v", rec)

			tx.Rollback()
		}
	}()

	// get last monster code
	lastMonsterCode, err := uMonster.monsterRepo.GetLastMonsterCode()
	if err != nil {
		return http.StatusInternalServerError, "failed to get last monster code", err
	}

	// mapping req create monster
	reqMonster := model.Monster{
		Name:              req.Name,
		MonsterCode:       lastMonsterCode + 1,
		MonsterCategoryId: req.MonsterCategoryId,
		Description:       req.Description,
		Length:            req.Length,
		Weight:            req.Weight,
		HP:                req.HP,
		Attack:            req.Attack,
		Defends:           req.Defends,
		Speed:             req.Speed,
		ImageName:         req.ImageName,
	}

	// create database transaction
	trx, resCode, err := uMonster.monsterRepo.Transaction()
	if err != nil {
		return resCode, "failed to create database transaction", err
	}
	tx = trx.Begin()

	// create monster
	monsterId, err := uMonster.monsterRepo.CreateMonster(tx, ctx, reqMonster)
	if err != nil {
		if err.(*pq.Error).Code == "23503" {
			return http.StatusBadRequest, "invalid monster category", err
		}
		return http.StatusInternalServerError, "failed to create monster", err
	}

	// mapping req create mapping monster and monster type
	var reqMonsterAndType []model.MappingMonsterAndTypes
	if req.MonsterTypes != nil {
		for i := 0; i < len(req.MonsterTypes); i++ {
			monsterAndType := model.MappingMonsterAndTypes{
				MonsterId:     monsterId,
				MonsterTypeId: req.MonsterTypes[i],
			}
			reqMonsterAndType = append(reqMonsterAndType, monsterAndType)
		}

		// create mapping monster and monster type
		err = uMonster.monsterRepo.CreateMappingMonsterAndType(tx, ctx, reqMonsterAndType)
		if err != nil {
			if err.(*pq.Error).Code == "23503" {
				return http.StatusBadRequest, "invalid monster or monster type data", err
			}
			return http.StatusInternalServerError, "failed to create monster type", err
		}
	}

	// commit database transaction
	err = tx.Commit().Error
	if err != nil {
		return http.StatusInternalServerError, "failed to commit database transaction", err
	}

	// mapping response data
	resCode = http.StatusCreated
	resMessage = "create monster successfully"
	err = nil

	return resCode, resMessage, err
}

// GetMonsterById is use case to get monster by id
func (uMonster *monsterUseCase) GetMonsterById(ctx context.Context, reqId string) (res model.GetDetailMonsterRes, resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	// query get params
	queryGetParams := map[string]interface{}{
		"preloadParams": map[string]interface{}{
			"MonsterCategory": true,
			"MonsterTypes":    true,
		},
	}

	// find monster by id
	resMonster, err := uMonster.monsterRepo.GetMonsterById(ctx, reqId, queryGetParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, http.StatusBadRequest, "monster not found", err
		}
		return res, http.StatusInternalServerError, "failed to get monster by id", err
	}

	// mapping response data
	res.ID = resMonster.ID
	res.MonsterCode = resMonster.MonsterCode
	res.Name = resMonster.Name
	res.MonsterCategory = resMonster.MonsterCategory
	res.MonsterTypes = resMonster.MonsterTypes
	res.Description = resMonster.Description
	res.Length = resMonster.Length
	res.Weight = resMonster.Weight
	res.HP = resMonster.HP
	res.Attack = resMonster.Attack
	res.Defends = resMonster.Defends
	res.Speed = resMonster.Speed
	res.IsCaught = resMonster.IsCaught
	res.ImageName = resMonster.ImageName
	res.ImageURL = fmt.Sprintf("%s/api/v1/monster/images/%s", uMonster.baseURL, resMonster.ImageName)

	return res, http.StatusOK, "get monster successfully", nil
}

// GetListMonster is use case to get list monster
func (uMonster *monsterUseCase) GetListMonster(ctx context.Context, queryReq model.MonsterQueryReq) (res []model.GetListMonsterRes, resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	// query get params
	queryGetParams := map[string]interface{}{
		"selectParams": []string{`monsters.id, monsters.monster_code, monsters.name, monsters.monster_category_id, monsters.image_name, monsters.created_at`},
		"whereParams": map[string]interface{}{
			"default": map[string]interface{}{},
			"in":      map[string]interface{}{},
		},
		"preloadParams": map[string]interface{}{
			"MonsterCategory": true,
			"MonsterTypes":    true,
		},
		"joinParams": map[string]interface{}{},
	}
	if queryReq.Name != "" {
		queryGetParams["whereParams"].(map[string]interface{})["default"].(map[string]interface{})["lower(monsters.name) LIKE lower(?)"] = "%" + queryReq.Name + "%"
	}
	if queryReq.MonsterTypeId != nil {
		queryGetParams["selectParams"] = []string{`DISTINCT monsters.id, monsters.monster_code, monsters.name, monsters.monster_category_id, monsters.image_name, monsters.created_at`}
		queryGetParams["joinParams"].(map[string]interface{})["INNER JOIN mapping_monster_and_types map ON map.monster_id = monsters.id"] = true
		queryGetParams["whereParams"].(map[string]interface{})["in"].(map[string]interface{})["map.monster_type_id IN (?)"] = queryReq.MonsterTypeId
	}
	if queryReq.IsCaught != "" {
		isCaughtBool, err := strconv.ParseBool(queryReq.IsCaught)
		if err != nil {
			return nil, http.StatusBadRequest, "is_caught format not valid", err
		}
		queryGetParams["whereParams"].(map[string]interface{})["default"].(map[string]interface{})["monsters.is_caught = ?"] = isCaughtBool
	}

	// find list monster
	resMonster, err := uMonster.monsterRepo.GetListMonster(ctx, queryReq, queryGetParams)
	if err != nil {
		return res, http.StatusInternalServerError, "failed to get list monster", err
	}

	// mapping response data
	for i := 0; i < len(resMonster); i++ {
		monster := model.GetListMonsterRes{
			ID:              resMonster[i].ID,
			MonsterCode:     resMonster[i].MonsterCode,
			Name:            resMonster[i].Name,
			MonsterCategory: resMonster[i].MonsterCategory,
			MonsterTypes:    resMonster[i].MonsterTypes,
			IsCaught:        resMonster[i].IsCaught,
			ImageName:       resMonster[i].ImageName,
			ImageURL:        fmt.Sprintf("%s/api/v1/monster/images/%s", uMonster.baseURL, resMonster[i].ImageName),
		}
		res = append(res, monster)
	}

	return res, http.StatusOK, "get all monster successfully", nil
}

// UpdateMonster is use case to update monster
func (uMonster *monsterUseCase) UpdateMonster(ctx context.Context, reqId string, req model.UpdateMonsterReq) (resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	var tx = &gorm.DB{}
	defer func() {
		if rec := recover(); rec != nil {
			// mapping response data
			resCode = http.StatusInternalServerError
			resMessage = "failed update monster"
			err = fmt.Errorf("%v", rec)

			tx.Rollback()
		}
	}()

	// query get params
	queryGetParams := map[string]interface{}{
		"selectParams": []string{`id, name, monster_category_id, description, length, weight, hp, attack, defends, speed, is_caught, image_name`},
		"preloadParams": map[string]interface{}{
			"MonsterTypes": true,
		},
	}

	// find monster by id
	resMonster, err := uMonster.monsterRepo.GetMonsterById(ctx, reqId, queryGetParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, "monster not found", err
		}
		return http.StatusInternalServerError, "failed to get monster by id", err
	}

	// query update params
	queryUpdateParams := map[string]interface{}{
		"value": map[string]interface{}{},
		"whereParams": map[string]interface{}{
			"default": map[string]interface{}{
				"id = ?": reqId,
			},
		},
	}

	// mapping req update monster
	if req.Name != "" {
		queryUpdateParams["value"].(map[string]interface{})["name"] = req.Name
	}
	if req.MonsterCategoryId != "" {
		queryUpdateParams["value"].(map[string]interface{})["monster_category_id"] = req.MonsterCategoryId
	}
	if req.Description != "" {
		queryUpdateParams["value"].(map[string]interface{})["description"] = req.Description
	}
	if req.Length != resMonster.Length {
		queryUpdateParams["value"].(map[string]interface{})["length"] = req.Length
	}
	if req.Weight != resMonster.Weight {
		queryUpdateParams["value"].(map[string]interface{})["weight"] = req.Weight
	}
	if req.HP != resMonster.HP {
		queryUpdateParams["value"].(map[string]interface{})["hp"] = req.HP
	}
	if req.Attack != resMonster.Attack {
		queryUpdateParams["value"].(map[string]interface{})["attack"] = req.Attack
	}
	if req.Defends != resMonster.Defends {
		queryUpdateParams["value"].(map[string]interface{})["defends"] = req.Defends
	}
	if req.Speed != resMonster.Speed {
		queryUpdateParams["value"].(map[string]interface{})["speed"] = req.Speed
	}
	if req.IsCaught != resMonster.IsCaught {
		queryUpdateParams["value"].(map[string]interface{})["is_caught"] = req.IsCaught
	}
	if req.ImageName != "" {
		queryUpdateParams["value"].(map[string]interface{})["image_name"] = req.ImageName
		err = uploader.DeleteImage(resMonster.ImageName)
		if err != nil {
			return http.StatusInternalServerError, "failed to update monster image", err
		}
	}

	// create database transaction
	trx, resCode, err := uMonster.monsterRepo.Transaction()
	if err != nil {
		return resCode, "failed to create database transaction", err
	}
	tx = trx.Begin()

	// update monster
	err = uMonster.monsterRepo.UpdateMonster(tx, ctx, queryUpdateParams)
	if err != nil {
		if err.(*pq.Error).Code == "23503" {
			return http.StatusBadRequest, "invalid monster category", err
		}
		return http.StatusInternalServerError, "failed to update monster", err
	}

	// mapping req update mapping monster and monster type
	var reqMonsterAndType []model.MappingMonsterAndTypes
	if req.MonsterTypes != nil {
		for i := 0; i < len(req.MonsterTypes); i++ {
			monsterAndType := model.MappingMonsterAndTypes{
				MonsterId:     reqId,
				MonsterTypeId: req.MonsterTypes[i],
			}
			reqMonsterAndType = append(reqMonsterAndType, monsterAndType)
		}

		// delete mapping monster and monster type
		if resMonster.MonsterTypes != nil {
			err = uMonster.monsterRepo.DeleteMappingMonsterAndType(tx, ctx, reqId)
			if err != nil {
				return http.StatusInternalServerError, "failed to update monster type", err
			}
		}

		// create mapping monster and monster type
		err = uMonster.monsterRepo.CreateMappingMonsterAndType(tx, ctx, reqMonsterAndType)
		if err != nil {
			if err.(*pq.Error).Code == "23503" {
				return http.StatusBadRequest, "invalid monster or monster type data", err
			}
			return http.StatusInternalServerError, "failed to update monster type", err
		}
	}

	// commit database transaction
	err = tx.Commit().Error
	if err != nil {
		return http.StatusInternalServerError, "failed to commit database transaction", err
	}

	// mapping response data
	resCode = http.StatusOK
	resMessage = "update monster successfully"
	err = nil

	return resCode, resMessage, err
}

// UpdateMonsterCaptured is use case to update monster captured mark
func (uMonster *monsterUseCase) UpdateMonsterCaptured(ctx context.Context, reqId string, req model.UpdateMonsterCapturedReq) (resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	var tx = &gorm.DB{}
	defer func() {
		if rec := recover(); rec != nil {
			// mapping response data
			resCode = http.StatusInternalServerError
			resMessage = "failed update monster captured mark"
			err = fmt.Errorf("%v", rec)

			tx.Rollback()
		}
	}()

	// query get params
	queryGetParams := map[string]interface{}{
		"selectParams": []string{`id, is_caught`},
	}

	// find monster by id
	resMonster, err := uMonster.monsterRepo.GetMonsterById(ctx, reqId, queryGetParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, "monster not found", err
		}
		return http.StatusInternalServerError, "failed to get monster by id", err
	}

	// query update params
	queryUpdateParams := map[string]interface{}{
		"value": map[string]interface{}{},
		"whereParams": map[string]interface{}{
			"default": map[string]interface{}{
				"id = ?": reqId,
			},
		},
	}

	// mapping req update monster
	if req.IsCaught != resMonster.IsCaught {
		queryUpdateParams["value"].(map[string]interface{})["is_caught"] = req.IsCaught
	}

	// create database transaction
	trx, resCode, err := uMonster.monsterRepo.Transaction()
	if err != nil {
		return resCode, "failed to create database transaction", err
	}
	tx = trx.Begin()

	// update monster
	err = uMonster.monsterRepo.UpdateMonster(tx, ctx, queryUpdateParams)
	if err != nil {
		return http.StatusInternalServerError, "failed to update monster captured mark", err
	}

	// commit database transaction
	err = tx.Commit().Error
	if err != nil {
		return http.StatusInternalServerError, "failed to commit database transaction", err
	}

	// mapping response data
	resCode = http.StatusOK
	resMessage = "update monster captured mark successfully"
	err = nil

	return resCode, resMessage, err
}

// DeleteMonster is use case to delete monster
func (uMonster *monsterUseCase) DeleteMonster(ctx context.Context, reqId string) (resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMonster.ctxTimeout)
	defer cancel()

	var tx = &gorm.DB{}
	defer func() {
		if rec := recover(); rec != nil {
			// mapping response data
			resCode = http.StatusInternalServerError
			resMessage = "failed delete monster"
			err = fmt.Errorf("%v", rec)

			tx.Rollback()
		}
	}()

	// query get params
	queryGetParams := map[string]interface{}{
		"selectParams": []string{`id, image_name`},
		"preloadParams": map[string]interface{}{
			"MonsterTypes": true,
		},
	}

	// find monster by id
	resMonster, err := uMonster.monsterRepo.GetMonsterById(ctx, reqId, queryGetParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, "monster not found", err
		}
		return http.StatusInternalServerError, "failed to get monster by id", err
	}

	// query delete params
	queryDeleteParams := map[string]interface{}{
		"whereParams": map[string]interface{}{
			"default": map[string]interface{}{
				"id = ?": reqId,
			},
		},
	}

	// create database transaction
	trx, resCode, err := uMonster.monsterRepo.Transaction()
	if err != nil {
		return resCode, "failed to create database transaction", err
	}
	tx = trx.Begin()

	// soft delete monster
	err = uMonster.monsterRepo.SoftDeleterMonster(tx, ctx, queryDeleteParams)
	if err != nil {
		return resCode, "failed to delete monster", err
	}

	// delete image
	if resMonster.ImageName != "" {
		err = uploader.DeleteImage(resMonster.ImageName)
		if err != nil {
			return http.StatusInternalServerError, "failed to delete monster image", err
		}
	}

	// delete mapping monster and monster type if exist
	if resMonster.MonsterTypes != nil {
		err = uMonster.monsterRepo.DeleteMappingMonsterAndType(tx, ctx, reqId)
		if err != nil {
			return http.StatusInternalServerError, "failed to update monster type", err
		}
	}

	// commit database transaction
	err = tx.Commit().Error
	if err != nil {
		return http.StatusInternalServerError, "failed to commit database transaction", err
	}

	// mapping response data
	resCode = http.StatusOK
	resMessage = "delete monster successfully"
	err = nil

	return resCode, resMessage, err
}
