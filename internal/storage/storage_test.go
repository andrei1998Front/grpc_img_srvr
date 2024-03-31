package storage

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
	"github.com/stretchr/testify/require"
)

type TestImgInfo struct {
	FileName string
	Size     uint32
}

func prepareTestListImages(path string) []*TestImgInfo {
	var listOfImages []*TestImgInfo

	err := filepath.WalkDir(path, func(pathImg string, d os.DirEntry, err error) error {
		ch := checkExtension(AllowedExtensions, filepath.Ext(d.Name()))
		cd := d.IsDir()

		if !cd && ch {
			info, err := d.Info()

			if err != nil {
				return err
			}

			imgInfo := &TestImgInfo{
				FileName: d.Name(),
				Size:     uint32(info.Size()),
			}

			listOfImages = append(listOfImages, imgInfo)
		}
		return nil
	})

	if err != nil {
		return []*TestImgInfo{}
	}

	return listOfImages
}

func Test_prepareListImages(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    []*TestImgInfo
		wantErr bool
		err     string
	}{
		{
			name:    "Success",
			path:    "../../imgs",
			want:    prepareTestListImages("../../imgs"),
			wantErr: false,
		},
		{
			name:    "Non-existent path",
			path:    "fdfdfsdfdsfg",
			wantErr: true,
			err:     ErrPathNotExists.Error(),
		},
		{
			name:    "Path is not directory",
			path:    "../../imgs/cat.jpg",
			wantErr: true,
			err:     ErrPathNotDirectory.Error(),
		},
	}
	for _, tt := range tests {
		got, err := prepareListImages(tt.path)

		fmt.Println(err)

		if !tt.wantErr && err == nil {
			require.Equal(t, len(tt.want), len(got))

			for i, v := range got {
				require.Equal(t, v.FileName, tt.want[i].FileName)
				require.Equal(t, v.Size, tt.want[i].Size)
			}
		} else if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}
	}
}

func TestDiskImageStorage_Upload(t *testing.T) {
	mockStorage := DiskImageStorage{
		Path: "./test_imgs_dir",
		ListImages: []*models.ImgInfo{
			{
				FileName: "dddd.jpeg",
				Path:     "./test_imgs_dir",
				Size:     uint32(33),
				CreateDt: time.Now(),
				UpdateDt: time.Now(),
			},
		},
	}

	type args struct {
		filename string
		data     bytes.Buffer
	}
	tests := []struct {
		name    string
		d       *DiskImageStorage
		args    args
		want    models.ImgInfo
		lenList int
		wantErr bool
		err     string
	}{
		{
			name: "Success",
			d:    &mockStorage,
			args: args{
				filename: "pupa.png",
				data:     bytes.Buffer{},
			},
			want: models.ImgInfo{
				FileName: "pupa.png",
				Path:     mockStorage.Path,
				Size:     uint32(0),
			},
			wantErr: false,
			lenList: 2,
		},
		{
			name: "Invalid extension",
			d:    &mockStorage,
			args: args{
				filename: "lupa.sql",
				data:     bytes.Buffer{},
			},
			want:    models.ImgInfo{},
			wantErr: true,
			lenList: 2,
			err:     ErrInvalidImgExtension.Error(),
		},
	}
	for _, tt := range tests {
		got, err := tt.d.Upload(tt.args.filename, tt.args.data)
		lenGot := len(tt.d.ListImages)

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}

		require.Equal(t, tt.want.FileName, got.FileName)
		require.Equal(t, tt.want.Path, got.Path)
		require.Equal(t, tt.lenList, lenGot)
	}
}

func TestDiskImageStorage_Download(t *testing.T) {
	mockStorage := DiskImageStorage{
		Path: "./test_imgs_dir",
		ListImages: []*models.ImgInfo{
			{
				FileName: "dddd.jpeg",
				Path:     "./test_imgs_dir",
				Size:     uint32(33),
				CreateDt: time.Now(),
				UpdateDt: time.Now(),
			},
		},
	}

	tests := []struct {
		name     string
		d        *DiskImageStorage
		filename string
		want     *models.ImgInfo
		wantErr  bool
		err      string
	}{
		{
			name:     "Success",
			d:        &mockStorage,
			filename: "dddd.jpeg",
			want: &models.ImgInfo{
				FileName: mockStorage.ListImages[0].FileName,
				Path:     mockStorage.Path,
				Size:     mockStorage.ListImages[0].Size,
			},
			wantErr: false,
		},
		{
			name:     "Non-existing img",
			d:        &mockStorage,
			filename: "ffff.jpeg",
			want:     &models.ImgInfo{},
			wantErr:  true,
			err:      ErrImageNotFound.Error(),
		},
	}
	for _, tt := range tests {
		got, err := tt.d.Download(tt.filename)

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}

		require.Equal(t, tt.want.FileName, got.FileName)
		require.Equal(t, tt.want.Path, got.Path)
		require.Equal(t, tt.want.Size, got.Size)
	}
}

func TestDiskImageStorage_ListOfImages(t *testing.T) {
	mockStorage := DiskImageStorage{
		Path: "./test_imgs_dir",
		ListImages: []*models.ImgInfo{
			{
				FileName: "dddd.jpeg",
				Path:     "./test_imgs_dir",
				Size:     uint32(33),
				CreateDt: time.Now(),
				UpdateDt: time.Now(),
			},
		},
	}

	tests := []struct {
		name    string
		d       *DiskImageStorage
		want    []*models.ImgInfo
		wantErr bool
		err     string
	}{
		{
			name:    "Success",
			d:       &mockStorage,
			want:    mockStorage.ListImages,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := tt.d.ListOfImages(context.Background())

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}

		require.Equal(t, tt.want, got)
	}
}

func Test_checkExtension(t *testing.T) {
	tests := []struct {
		name    string
		ext     string
		listExt []string
		want    bool
	}{
		{
			name:    "True",
			ext:     ".jpeg",
			listExt: AllowedExtensions,
			want:    true,
		},
		{
			name:    "False",
			ext:     ".sql",
			listExt: AllowedExtensions,
			want:    false,
		},
	}
	for _, tt := range tests {
		got := checkExtension(AllowedExtensions, tt.ext)

		require.Equal(t, tt.want, got)
	}
}

func TestDiskImageStorage_CheckExists(t *testing.T) {
	mockStorage := DiskImageStorage{
		Path: "./test_imgs_dir",
		ListImages: []*models.ImgInfo{
			{
				FileName: "dddd.jpeg",
				Path:     "./test_imgs_dir",
				Size:     uint32(33),
				CreateDt: time.Now(),
				UpdateDt: time.Now(),
			},
		},
	}

	tests := []struct {
		name     string
		filename string
		d        *DiskImageStorage
		want     bool
	}{
		{
			name:     "True",
			filename: "dddd.jpeg",
			d:        &mockStorage,
			want:     true,
		},
		{
			name:     "False",
			filename: "ddddf.jpeg",
			d:        &mockStorage,
			want:     false,
		},
	}
	for _, tt := range tests {
		got := tt.d.CheckExists(tt.filename)

		require.Equal(t, tt.want, got)
	}
}
