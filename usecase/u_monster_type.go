package usecase

import (
	"context"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/repository"
	"net/http"
	"time"
)

// MTypeUseCaseInterface is
type MTypeUseCaseInterface interface {
	GetAllMonsterType(ctx context.Context) (res []model.MonsterType, resCode int, resMessage string, err error)
}

type mTypeUseCase struct {
	ctxTimeout time.Duration
	mTypeRepo  repository.MTypeRepositoryInterface
}

func NewMTypeUseCase(ctxTimeout time.Duration, mTypeRepo repository.MTypeRepositoryInterface) MTypeUseCaseInterface {
	return &mTypeUseCase{
		ctxTimeout: ctxTimeout,
		mTypeRepo:  mTypeRepo,
	}
}

// GetAllMonsterType is use case to get all monster type
func (uMType *mTypeUseCase) GetAllMonsterType(ctx context.Context) (res []model.MonsterType, resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMType.ctxTimeout)
	defer cancel()

	// find all monster type
	res, err = uMType.mTypeRepo.GetAllMonsterType(ctx, []string{`id`, `name`})
	if err != nil {
		return res, http.StatusInternalServerError, "failed to get all monster type", err
	}

	return res, http.StatusOK, "get all monster type successfully", nil
}
