package storage

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
)

var allowedExtensions []string = []string{".jpeg", ".jpg", ".png", ".svg"}

type DiskImageStorage struct {
	Path       string
	mutex      sync.RWMutex
	ListImages []*models.ImgInfo
}

func New(path string) *DiskImageStorage {

	listOfImages, err := prepareListImages(path)

	if err != nil {
		panic("storage initialization failure")
	}

	return &DiskImageStorage{
		Path:       path,
		ListImages: listOfImages,
	}
}

func checkExtension(listExtensions []string, currentExt string) bool {
	for _, v := range listExtensions {
		if currentExt == v {
			return true
		}
	}

	return false
}

func prepareListImages(path string) ([]*models.ImgInfo, error) {
	const op = "storage.prepartListImages"
	var listOfImages []*models.ImgInfo

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		ch := checkExtension(allowedExtensions, filepath.Ext(d.Name()))
		cd := d.IsDir()

		if !cd && ch {
			info, err := d.Info()

			if err != nil {
				return err
			}

			imgInfo := &models.ImgInfo{
				FileName: d.Name(),
				Size:     uint32(info.Size()),
				CreateDt: info.ModTime(),
				UpdateDt: info.ModTime(),
			}

			listOfImages = append(listOfImages, imgInfo)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return listOfImages, nil
}

func (d *DiskImageStorage) Upload(filename string, data bytes.Buffer) (models.ImgInfo, error) {
	const op string = "storage.Upload"

	imagePath := fmt.Sprintf("%s/%s", d.Path, filename)

	file, err := os.Create(imagePath)

	if err != nil {
		return models.ImgInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = data.WriteTo(file)
	if err != nil {
		return models.ImgInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	st, err := os.Stat(d.Path + "/" + filename)

	if err != nil {
		return models.ImgInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	imgInfo := models.ImgInfo{
		FileName: filename,
		Size:     uint32(st.Size()),
		CreateDt: st.ModTime(),
		UpdateDt: st.ModTime(),
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.ListImages = append(d.ListImages, &imgInfo)

	return imgInfo, nil
}

func (d *DiskImageStorage) CheckExists(filename string) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	for _, v := range d.ListImages {
		if filename == v.FileName {
			return true
		}
	}

	return false
}
