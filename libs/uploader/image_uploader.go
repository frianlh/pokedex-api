package uploader

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// SaveImage is
func SaveImage(ctx *fiber.Ctx, imageFile *multipart.FileHeader) (imageName string, err error) {
	path := "./images"

	// make directory, if not exist
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// save image
	newImageName := fmt.Sprintf("monster_%d%s", time.Now().Unix(), filepath.Ext(imageFile.Filename))
	err = ctx.SaveFile(imageFile, fmt.Sprintf("%s/%s", path, newImageName))
	if err != nil {
		return "", err
	}

	return newImageName, nil
}

// DeleteImage is
func DeleteImage(imageName string) (err error) {
	path := "./images"

	// delete image from directory
	err = os.Remove(fmt.Sprintf("%s/%s", path, imageName))
	if err != nil {
		return err
	}

	return nil
}
