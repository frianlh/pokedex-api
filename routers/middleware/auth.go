package middleware

import (
	"encoding/json"
	"github.com/frianlh/pokedex-api/libs/encrypt"
	"github.com/frianlh/pokedex-api/libs/response"
	"github.com/frianlh/pokedex-api/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// AuthMiddleware is function fo authentication middleware
func AuthMiddleware(jwtKey, menuPermission string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// token validation
		claims, err := encrypt.Parse(ctx.Get("Authorization"), jwtKey)
		if err != nil {
			return response.ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err.Error())
		}

		// get permission
		var permissions []model.Permission
		dataJSON, err := json.Marshal(claims["permission"])
		if err != nil {
			return response.ErrorRes(ctx, http.StatusInternalServerError, "unauthorized", err.Error())
		}
		if err = json.Unmarshal(dataJSON, &permissions); err != nil {
			return response.ErrorRes(ctx, http.StatusInternalServerError, "unauthorized", err.Error())
		}
		isAccessExist := false
		for i := 0; i < len(permissions); i++ {
			if permissions[i].Name == menuPermission {
				isAccessExist = true
				break
			}
		}
		if !isAccessExist {
			return response.ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", "unauthorized")
		}

		return ctx.Next()
	}
}
