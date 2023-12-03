package delivery

import (
	"github.com/frianlh/pokedex-api/libs/response"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type authHandler struct {
	authUseCase usecase.AuthUseCaseInterface
}

func NewAuthHandler(authUseCase usecase.AuthUseCaseInterface) *authHandler {
	return &authHandler{
		authUseCase: authUseCase,
	}
}

// Login is handler for user login
func (hAuth *authHandler) Login(ctx *fiber.Ctx) error {
	var req model.LoginReq

	// binding request body to struct
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "failed to binds the request body", err.Error())
	}

	// login
	res, resCode, resMessage, err := hAuth.authUseCase.Login(ctx.Context(), req)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", res)
}
