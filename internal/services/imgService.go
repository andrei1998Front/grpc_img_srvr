package imgService

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
)

var (
	ErrImageExists = errors.New("image already exists")
)

type ImgService struct {
	log                *slog.Logger
	imgUploader        ImgUploader
	imgDownloader      ImgDownloader
	listOfImagesGetter ListOfImagesGetter
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=ImgUploader
type ImgUploader interface {
	Upload(
		filename string,
		chunk bytes.Buffer,
	) (models.ImgInfo, error)
	CheckExists(filename string) bool
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=ImgDownloader
type ImgDownloader interface {
	Download(
		filename string,
	) (*models.ImgInfo, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=ListOfImagesGetter
type ListOfImagesGetter interface {
	ListOfImages(ctx context.Context) ([]*models.ImgInfo, error)
}

func New(
	log *slog.Logger,
	imgUploader ImgUploader,
	imgDownloader ImgDownloader,
	listOfImagesGetter ListOfImagesGetter,
) *ImgService {
	return &ImgService{
		log:                log,
		imgUploader:        imgUploader,
		imgDownloader:      imgDownloader,
		listOfImagesGetter: listOfImagesGetter,
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
		log.Error("imagez "+filename+" upload failed", err)
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

func (s ImgService) ListOfImages(ctx context.Context) (string, error) {
	const op = "ImageService.ListOfImages"

	log := s.log.With(slog.String(op, "op"))

	log.Info("start receiving list of images from storage")

	lof, err := s.listOfImagesGetter.ListOfImages(ctx)

	if err != nil {
		log.Error("receive list of images failed", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("the list of images was received from the storage successfully")

	return lofToStr(lof), nil
}

func lofToStr(infoList []*models.ImgInfo) string {
	var strSlice []string

	for _, info := range infoList {
		infoStr := imgInfoToStr(info)

		strSlice = append(strSlice, infoStr)
	}

	return strings.Join(strSlice, "\n")
}

func imgInfoToStr(infoItem *models.ImgInfo) string {
	if infoItem.FileName == "" {
		return ""
	}

	createDt := infoItem.CreateDt.Format("2006-01-02 15:04:05")
	updateDt := infoItem.UpdateDt.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%s | %s | %s", infoItem.FileName, createDt, updateDt)
}
