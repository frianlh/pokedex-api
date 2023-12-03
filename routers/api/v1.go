package api

import (
	"github.com/frianlh/pokedex-api/configs"
	"github.com/frianlh/pokedex-api/delivery"
	"github.com/frianlh/pokedex-api/repository"
	"github.com/frianlh/pokedex-api/routers/middleware"
	"github.com/frianlh/pokedex-api/usecase"
	"github.com/gofiber/fiber/v2"
)

// V1Route is routers for dashboard version 1
func V1Route(route fiber.Router, config *configs.Config) {
	// repository
	rUser := repository.NewUserRepository(config.PostgresConfig.DbConn)
	rMCategory := repository.NewMonsterCategory(config.PostgresConfig.DbConn)
	rMType := repository.NewMTypeRepository(config.PostgresConfig.DbConn)
	rMonster := repository.NewMonsterRepository(config.PostgresConfig.DbConn)

	// use case
	uAuth := usecase.NewAuthUseCase(config.TimeoutCtx, config.JWTKey, rUser)
	uMCategory := usecase.NewMCategoryUseCase(config.TimeoutCtx, rMCategory)
	uMType := usecase.NewMTypeUseCase(config.TimeoutCtx, rMType)
	uMonster := usecase.NewMonsterUseCase(config.TimeoutCtx, config.BaseURL, rMonster)

	// delivery
	hAuth := delivery.NewAuthHandler(uAuth)
	hMCategory := delivery.NewMCategoryHandler(uMCategory)
	hMType := delivery.NewMTypeHandler(uMType)
	hMonster := delivery.NewMonsterHandler(uMonster)

	// route group
	// auth group
	auth := route.Group("/auth")
	{
		auth.Post("/login", hAuth.Login)
	}

	// monster category group
	mCategory := route.Group("/monster-category")
	{
		mCategory.Get("", hMCategory.GetAllMonsterCategory)
	}

	// monster type group
	mType := route.Group("/monster-type")
	{
		mType.Get("", hMType.GetAllMonsterType)
	}

	// monster group
	monster := route.Group("/monster")
	{
		monster.Post("", middleware.AuthMiddleware(config.JWTKey, "write_monster"), hMonster.CreateMonster)
		monster.Get("/:id", hMonster.GetMonsterById)
		monster.Get("", hMonster.GetListMonster)
		monster.Put("/:id", middleware.AuthMiddleware(config.JWTKey, "update_monster"), hMonster.UpdateMonster)
		monster.Put("captured/:id", hMonster.UpdateMonsterCaptured)
		monster.Delete("/:id", middleware.AuthMiddleware(config.JWTKey, "delete_monster"), hMonster.DeleteMonster)
		monster.Static("/images", "./images")
	}
}
