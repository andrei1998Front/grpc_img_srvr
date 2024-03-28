package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

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
