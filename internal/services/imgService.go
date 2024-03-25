package imgService

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
)

var (
	ErrImageExists = errors.New("image already exists")
)

type ImgService struct {
	log           *slog.Logger
	imgUploader   ImgUploader
	imgDownloader ImgDownloader
}

type ImgUploader interface {
	Upload(
		filename string,
		chunk bytes.Buffer,
	) (models.ImgInfo, error)
	CheckExists(filename string) bool
}

type ImgDownloader interface {
	Download(
		filename string,
	) (*models.ImgInfo, error)
}

func New(
	log *slog.Logger,
	imgUploader ImgUploader,
	imgDownloader ImgDownloader,
) *ImgService {
	return &ImgService{
		log:           log,
		imgUploader:   imgUploader,
		imgDownloader: imgDownloader,
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

func (s ImgService) Download(
	filename string,
) (*models.ImgInfo, error) {
	const op = "ImageService.Upload"

	log := s.log.With(
		slog.String("op", op),
		slog.String("filename", filename),
	)

	log.Info("start receiving file " + filename + " from storage")

	imgInfo, err := s.imgDownloader.Download(filename)

	if err != nil {
		log.Error("image "+filename+" download failed", err)
		return &models.ImgInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("the file " + filename + " was received from the storage successfully")

	return imgInfo, nil
}
