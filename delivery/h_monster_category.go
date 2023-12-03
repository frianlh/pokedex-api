package delivery

import (
	"github.com/frianlh/pokedex-api/libs/response"
	"github.com/frianlh/pokedex-api/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type mCategoryHandler struct {
	mCategoryUseCase usecase.MCategoryUseCaseInterface
}

func NewMCategoryHandler(mCategoryUseCase usecase.MCategoryUseCaseInterface) *mCategoryHandler {
	return &mCategoryHandler{
		mCategoryUseCase: mCategoryUseCase,
	}
}

// GetAllMonsterCategory is handler to get all monster category
func (hMCategory *mCategoryHandler) GetAllMonsterCategory(ctx *fiber.Ctx) error {
	// find all monster category
	res, resCode, resMessage, err := hMCategory.mCategoryUseCase.GetAllMonsterCategory(ctx.Context())
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", res)
}
