package storage

import (
	"bytes"
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
		ch := checkExtension(allowedExtensions, filepath.Ext(d.Name()))
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
	}
	for _, tt := range tests {
		got, err := tt.d.Upload(tt.args.filename, tt.args.data)

		t.Log(err)
		require.Equal(t, tt.want.FileName, got.FileName)
	}
}
