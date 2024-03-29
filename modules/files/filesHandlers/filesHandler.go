package filesHandlers

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"

	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/entities"
	"github.com/Japanisnmm/GoBackend101/modules/files"
	"github.com/Japanisnmm/GoBackend101/modules/files/filesUsecases"
	"github.com/Japanisnmm/GoBackend101/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type filesHandlersErrCode string
const (
	uploadErr filesHandlersErrCode = "files-001"
)


type IFileHandlers interface{
	UploadFiles(c *fiber.Ctx) error
}

type filesHandler struct {
	cfg config.IConfig
	filesUsecase filesUsecases.IFilesUsecase
}

func FileHandler(cfg config.IConfig,filesUsecase filesUsecases.IFilesUsecase) IFileHandlers {
	return &filesHandler{
		cfg: cfg,
		filesUsecase: filesUsecase,
	}
}

func (h *filesHandler) UploadFiles(c *fiber.Ctx) error {
	req := make([]*files.FileReq, 0)

	form, err := c.MultipartForm()
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(uploadErr),
			err.Error(),
		).Res()
	}
	filesReq := form.File["files"]
	destination := c.FormValue("destination")

	// Files ext validation
	extMap := map[string]string{
		"png":  "png",
		"jpg":  "jpg",
		"jpeg": "jpeg",
	}
	for _, file := range filesReq {
		ext := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
		if extMap[ext] != ext || extMap[ext] == "" {
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(uploadErr),
				"extension is not acceptable",
			).Res()
		}

		if file.Size > int64(h.cfg.App().FileLimit()) {
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(uploadErr),
				fmt.Sprintf("file size must less than %d MiB", int(math.Ceil(float64(h.cfg.App().FileLimit())/math.Pow(1024, 2)))),
			).Res()
		}

		filename := utils.RandFileName(ext)
		req = append(req, &files.FileReq{
			File:        file,
			Destination: destination + "/" + filename,
			FileName:    filename,
			Extension:   ext,
		})
	}

	res, err := h.filesUsecase.UploadToGCP(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(uploadErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(fiber.StatusCreated, res).Res()
}