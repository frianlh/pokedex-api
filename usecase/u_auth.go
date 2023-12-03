package usecase

import (
	"context"
	"errors"
	"github.com/frianlh/pokedex-api/libs/encrypt"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/repository"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// AuthUseCaseInterface is
type AuthUseCaseInterface interface {
	Login(ctx context.Context, req model.LoginReq) (res model.LoginRes, resCode int, resMessage string, err error)
}

type authUseCase struct {
	ctxTimeout time.Duration
	jwtKey     string
	userRepo   repository.UserRepositoryInterface
}

func NewAuthUseCase(ctxTimeout time.Duration, jwtKey string, userRepo repository.UserRepositoryInterface) AuthUseCaseInterface {
	return &authUseCase{
		ctxTimeout: ctxTimeout,
		jwtKey:     jwtKey,
		userRepo:   userRepo,
	}
}

// Login is use case for user login
func (uAuth *authUseCase) Login(ctx context.Context, req model.LoginReq) (res model.LoginRes, resCode int, resMessage string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uAuth.ctxTimeout)
	defer cancel()

	// query get params
	queryGetParams := map[string]interface{}{
		"selectParams": []string{"id", "encrypted_password", "role_id"},
		"whereParams": map[string]interface{}{
			"default": map[string]interface{}{
				"email = ?": req.Email,
			},
		},
		"preloadParams": map[string]interface{}{
			"Role":             true,
			"Role.Permissions": true,
		},
	}

	// find user by email
	resUser, err := uAuth.userRepo.GetUserByParams(ctx, queryGetParams)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, http.StatusBadRequest, "email or password is incorrect", err
		}
		return res, http.StatusInternalServerError, "failed to get user", err
	}

	// password comparison
	err = encrypt.CompareHashAndPassword(&resUser.EncryptedPassword, &req.Password)
	if err != nil {
		return res, http.StatusBadRequest, "email or password is incorrect", err
	}

	// generate jwt
	authUser := jwt.MapClaims{
		"id":         resUser.ID,
		"role_id":    resUser.RoleId,
		"permission": resUser.Role.Permissions,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := encrypt.NewWithClaims(authUser, uAuth.jwtKey)
	res.Token = token

	return res, http.StatusOK, "login successfully", nil
}
