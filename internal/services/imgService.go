package imgService

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
)

var (
	ErrImageExists   = errors.New("image already exists")
	ErrImageNotFound = errors.New("image not found")
)

type ImgService struct {
	log         *slog.Logger
	imgUploader ImgUploader
}

type ImgUploader interface {
	Upload(
		filename string,
		chunk bytes.Buffer,
	) (models.ImgInfo, error)
	CheckExists(filename string) bool
}

func New(
	log *slog.Logger,
	imgUploader ImgUploader,
) *ImgService {
	return &ImgService{
		log:         log,
		imgUploader: imgUploader,
	}
}

func (s ImgService) Upload(
	filename string,
	chunk bytes.Buffer,
) (models.ImgInfo, error) {
	const op = "ImageService.Upload"

	log := s.log.With(
		slog.String("op", op),
		slog.String("filename", filename),
	)

	if s.imgUploader.CheckExists(filename) {
		log.Error(ErrImageExists.Error())
		return models.ImgInfo{}, ErrImageExists
	}

	log.Info("upload image " + filename)

	imgInfo, err := s.imgUploader.Upload(filename, chunk)

	if err != nil {
		log.Error("image "+filename+" upload failed", err)
		return models.ImgInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("image " + filename + " upload was successful")

	return imgInfo, nil
}
