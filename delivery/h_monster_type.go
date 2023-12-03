package delivery

import (
	"github.com/frianlh/pokedex-api/libs/response"
	"github.com/frianlh/pokedex-api/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type mTypeHandler struct {
	mTypeUseCase usecase.MTypeUseCaseInterface
}

func NewMTypeHandler(mTypeUseCase usecase.MTypeUseCaseInterface) *mTypeHandler {
	return &mTypeHandler{
		mTypeUseCase: mTypeUseCase,
	}
}

// GetAllMonsterType is handler to get all monster type
func (hMType *mTypeHandler) GetAllMonsterType(ctx *fiber.Ctx) error {
	// find all monster type
	res, resCode, resMessage, err := hMType.mTypeUseCase.GetAllMonsterType(ctx.Context())
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", res)
}
