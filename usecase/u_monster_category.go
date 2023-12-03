package usecase

import (
	"context"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/repository"
	"net/http"
	"time"
)

// MCategoryUseCaseInterface is
type MCategoryUseCaseInterface interface {
	GetAllMonsterCategory(ctx context.Context) (res []model.MonsterCategory, resCode int, resMessage string, err error)
}

type mCategoryUseCase struct {
	ctxTimeout    time.Duration
	mCategoryRepo repository.MCategoryRepositoryInterface
}

func NewMCategoryUseCase(ctxTimeout time.Duration, mCategoryRepo repository.MCategoryRepositoryInterface) MCategoryUseCaseInterface {
	return &mCategoryUseCase{
		ctxTimeout:    ctxTimeout,
		mCategoryRepo: mCategoryRepo,
	}
}

// GetAllMonsterCategory is use case to get all monster category
func (uMCategory *mCategoryUseCase) GetAllMonsterCategory(ctx context.Context) (res []model.MonsterCategory, resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uMCategory.ctxTimeout)
	defer cancel()

	// find all monster category
	res, err = uMCategory.mCategoryRepo.GetAllMonsterCategory(ctx, []string{`id`, `name`})
	if err != nil {
		return res, http.StatusInternalServerError, "failed to get all monster category", err
	}

	return res, http.StatusOK, "get all monster category successfully", nil
}
