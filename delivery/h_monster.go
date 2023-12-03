package delivery

import (
	"fmt"
	"github.com/frianlh/pokedex-api/libs/form"
	"github.com/frianlh/pokedex-api/libs/response"
	"github.com/frianlh/pokedex-api/libs/uploader"
	"github.com/frianlh/pokedex-api/libs/validator"
	"github.com/frianlh/pokedex-api/model"
	"github.com/frianlh/pokedex-api/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type monsterHandler struct {
	formatFile     map[string]bool
	monsterUseCase usecase.MonsterUseCaseInterface
}

func NewMonsterHandler(monsterUseCase usecase.MonsterUseCaseInterface) *monsterHandler {
	return &monsterHandler{
		formatFile: map[string]bool{
			".png":  true,
			".jpeg": true,
			".jpg":  true,
		},
		monsterUseCase: monsterUseCase,
	}
}

// CreateMonster is handler to create monster
func (hMonster *monsterHandler) CreateMonster(ctx *fiber.Ctx) error {
	var req model.CreateMonsterReq

	// binding request body to struct
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "failed to binds the request body", err.Error())
	}

	// validate request body
	// struct validation
	err = validator.ValidateStruct(&req)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "data input is invalid", err.Error())
	}

	// get monster type
	monsterTypeId, err := hMonster.getMonsterTypeId(ctx)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster type must be uuid", err.Error())
	}
	if len(monsterTypeId) == 0 {
		return response.ErrorRes(ctx, http.StatusBadRequest, "data input is invalid", "monster type is required")
	}
	req.MonsterTypes = monsterTypeId

	// get image file
	imageFile, err := ctx.FormFile("image")
	if err == nil {
		// image validation
		if imageFile == nil {
			return response.ErrorRes(ctx, http.StatusBadRequest, "data input is invalid", "image is required")
		}
		isValid, resMessage := hMonster.fileUploadedValidation(imageFile)
		if !isValid {
			return response.ErrorRes(ctx, http.StatusBadRequest, resMessage, err.Error())
		}

		// save image
		imageName, err := uploader.SaveImage(ctx, imageFile)
		if err != nil {
			return response.ErrorRes(ctx, http.StatusInternalServerError, "failed to save image", err.Error())
		}
		req.ImageName = imageName
	}

	// create monster
	resCode, resMessage, err := hMonster.monsterUseCase.CreateMonster(ctx.Context(), req)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusCreated, resMessage, "", nil)
}

// GetMonsterById is handler to get monster by id
func (hMonster *monsterHandler) GetMonsterById(ctx *fiber.Ctx) error {
	id := form.SQLInjector(ctx.Params("id"))
	_, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster id not valid", err.Error())
	}

	// find monster by id
	res, resCode, resMessage, err := hMonster.monsterUseCase.GetMonsterById(ctx.Context(), id)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", res)
}

// GetListMonster is handler to get list monster
func (hMonster *monsterHandler) GetListMonster(ctx *fiber.Ctx) error {
	sortBy := form.SQLInjector(ctx.Query("sort_by", "created_at"))
	orderBy := form.SQLInjector(strings.ToUpper(ctx.Query("order_by", "ASC")))
	name := form.SQLInjector(ctx.Query("name", ""))
	isCaught := form.SQLInjector(ctx.Query("is_caught", ""))

	// get monster type
	monsterTypeId, err := hMonster.getMonsterTypeId(ctx)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster type must be uuid", err.Error())
	}

	// query params request
	queryParams := model.MonsterQueryReq{
		SortBy:        sortBy,
		OrderBy:       orderBy,
		Name:          name,
		MonsterTypeId: monsterTypeId,
		IsCaught:      isCaught,
	}

	// find list monster
	res, resCode, resMessage, err := hMonster.monsterUseCase.GetListMonster(ctx.Context(), queryParams)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", res)
}

// UpdateMonster is handler to update monster
func (hMonster *monsterHandler) UpdateMonster(ctx *fiber.Ctx) error {
	var req model.UpdateMonsterReq

	id := form.SQLInjector(ctx.Params("id"))
	_, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster id not valid", err.Error())
	}

	// binding request body to struct
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "failed to binds the request body", err.Error())
	}

	// validate request body
	// get monster type
	monsterTypeId, err := hMonster.getMonsterTypeId(ctx)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster type must be uuid", err.Error())
	}
	req.MonsterTypes = monsterTypeId

	// get image file
	imageFile, err := ctx.FormFile("image")
	if err == nil {
		// image validation
		isValid, resMessage := hMonster.fileUploadedValidation(imageFile)
		if !isValid {
			return response.ErrorRes(ctx, http.StatusBadRequest, resMessage, err.Error())
		}

		// save image
		imageName, err := uploader.SaveImage(ctx, imageFile)
		if err != nil {
			return response.ErrorRes(ctx, http.StatusInternalServerError, "failed to save image", err.Error())
		}
		req.ImageName = imageName
	}

	// update monster
	resCode, resMessage, err := hMonster.monsterUseCase.UpdateMonster(ctx.Context(), id, req)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", nil)
}

// UpdateMonsterCaptured is handler to update monster captured mark
func (hMonster *monsterHandler) UpdateMonsterCaptured(ctx *fiber.Ctx) error {
	var req model.UpdateMonsterCapturedReq

	id := form.SQLInjector(ctx.Params("id"))
	_, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster id not valid", err.Error())
	}

	// binding request body to struct
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "failed to binds the request body", err.Error())
	}

	// update monster
	resCode, resMessage, err := hMonster.monsterUseCase.UpdateMonsterCaptured(ctx.Context(), id, req)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", nil)
}

// DeleteMonster is handler to delete monster
func (hMonster *monsterHandler) DeleteMonster(ctx *fiber.Ctx) error {
	id := form.SQLInjector(ctx.Params("id"))
	_, err := uuid.Parse(id)
	if err != nil {
		return response.ErrorRes(ctx, http.StatusBadRequest, "monster id not valid", err.Error())
	}

	// delete monster
	resCode, resMessage, err := hMonster.monsterUseCase.DeleteMonster(ctx.Context(), id)
	if err != nil {
		return response.ErrorRes(ctx, resCode, resMessage, err.Error())
	}

	return response.SuccessRes(ctx, http.StatusOK, resMessage, "", nil)
}

// getMonsterTypeId is
func (hMonster *monsterHandler) getMonsterTypeId(ctx *fiber.Ctx) (monsterType []string, err error) {
	loopCheck := true
	for i := 0; i < 7; i++ {
		typeId := ctx.FormValue(fmt.Sprintf("monster_type_id[%d]", i), "")
		if typeId != "" {
			_, err = uuid.Parse(typeId)
			if err != nil {
				return nil, err
			}
			monsterType = append(monsterType, typeId)
		} else {
			loopCheck = false
		}
		if !loopCheck {
			break
		}
	}

	return monsterType, nil
}

// fileUploadedValidation is
func (hMonster *monsterHandler) fileUploadedValidation(fileHeader *multipart.FileHeader) (isValid bool, resMessage string) {
	maxPartSize := int64(10 * 1024 * 1024)
	size := fileHeader.Size
	extension := filepath.Ext(fileHeader.Filename)

	if size > maxPartSize {
		return false, "file cannot exceed 10 MB"
	}
	if !hMonster.formatFile[extension] {
		return false, "file format must be .png, .jpg, or .jpeg"
	}

	return true, ""
}
