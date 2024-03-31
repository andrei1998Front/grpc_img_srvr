package imgService

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
	"github.com/andrei1998Front/grpc_img_srvr/internal/lib/slogdiscard"
	"github.com/andrei1998Front/grpc_img_srvr/internal/services/mocks"
	"github.com/andrei1998Front/grpc_img_srvr/internal/storage"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestImgService_Upload(t *testing.T) {
	type args struct {
		filename string
		chunk    bytes.Buffer
	}

	tests := []struct {
		name          string
		args          args
		want          models.ImgInfo
		wantErr       bool
		err           string
		mockCEoutput  bool
		mockUploadErr error
	}{
		{
			name: "Success",
			args: args{
				filename: "random.jpeg",
				chunk:    bytes.Buffer{},
			},
			want:         models.ImgInfo{FileName: "random.jpeg"},
			mockCEoutput: false,
			wantErr:      false,
		},
		{
			name: "Existing image",
			args: args{
				filename: "existing_image.jpeg",
				chunk:    bytes.Buffer{},
			},
			want:         models.ImgInfo{},
			err:          ErrImageExists.Error(),
			mockCEoutput: true,
			wantErr:      true,
		},
		{
			name: "Invalid extension",
			args: args{
				filename: "non_existing_image.sql",
				chunk:    bytes.Buffer{},
			},
			want:         models.ImgInfo{},
			err:          storage.ErrInvalidImgExtension.Error(),
			mockCEoutput: false,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		mk := mocks.NewImgUploader(t)
		mk.On("CheckExists", tt.args.filename).Return(tt.mockCEoutput).Once()

		if !tt.mockCEoutput {
			mk.On("Upload", tt.args.filename, tt.args.chunk).Return(tt.want, nil).Once()
		}
		s := New(
			slogdiscard.NewDiscardLogger(),
			mk,
			&mocks.ImgDownloader{},
			&mocks.ListOfImagesGetter{},
		)

		got, err := s.Upload(tt.args.filename, tt.args.chunk)

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}
		require.Equal(t, tt.want, got)
	}
}

func TestImgService_Download(t *testing.T) {
	tests := []struct {
		name              string
		filename          string
		want              *models.ImgInfo
		wantErr           bool
		err               string
		mockDownloadError error
	}{
		{
			name:     "Success",
			filename: "random.jpeg",
			want:     &models.ImgInfo{FileName: "random.jpeg"},
			wantErr:  false,
		},
		{
			name:              "Non-existing img",
			filename:          "non-existing.jpeg",
			want:              &models.ImgInfo{},
			wantErr:           true,
			err:               ErrImageExists.Error(),
			mockDownloadError: ErrImageExists,
		},
	}

	for _, tt := range tests {
		mk := mocks.NewImgDownloader(t)
		mk.On("Download", tt.filename).Return(tt.want, tt.mockDownloadError).Once()

		s := New(
			slogdiscard.NewDiscardLogger(),
			&mocks.ImgUploader{},
			mk,
			&mocks.ListOfImagesGetter{},
		)

		got, err := s.Download(tt.filename)

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}
		require.Equal(t, tt.want, got)
	}
}

func TestImgService_ListOfImages(t *testing.T) {
	tests := []struct {
		name                   string
		want                   string
		wantErr                bool
		err                    string
		mockListOfImagesOutput []*models.ImgInfo
		mockListOfImagesError  error
	}{
		{
			name: "Success. One item of list",
			want: "random.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05",
			mockListOfImagesOutput: []*models.ImgInfo{
				{
					FileName: "random.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
			},
		},
		{
			name: "Success. Multiple list items",
			want: "random_1.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05\nrandom_2.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05",
			mockListOfImagesOutput: []*models.ImgInfo{
				{
					FileName: "random_1.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
				{
					FileName: "random_2.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
			},
		},
		{
			name:                   "Success. Empty list",
			want:                   "",
			mockListOfImagesOutput: []*models.ImgInfo{},
		},
	}

	for _, tt := range tests {
		mk := mocks.NewListOfImagesGetter(t)
		mk.On("ListOfImages", mock.Anything).Return(tt.mockListOfImagesOutput, tt.mockListOfImagesError).Once()

		s := New(
			slogdiscard.NewDiscardLogger(),
			&mocks.ImgUploader{},
			&mocks.ImgDownloader{},
			mk,
		)

		got, err := s.ListOfImages(context.Background())

		if !tt.wantErr && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		} else if tt.wantErr && err != nil {
			require.ErrorContains(t, err, tt.err)
		}
		require.Equal(t, tt.want, got)
	}
}

func Test_lofToStr(t *testing.T) {
	tests := []struct {
		name     string
		infoList []*models.ImgInfo
		want     string
	}{
		{
			name: "Success. One item of list",
			want: "random.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05",
			infoList: []*models.ImgInfo{
				{
					FileName: "random.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
			},
		},
		{
			name: "Success. Multiple list items",
			want: "random_1.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05\nrandom_2.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05",
			infoList: []*models.ImgInfo{
				{
					FileName: "random_1.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
				{
					FileName: "random_2.jpeg",
					CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
					UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				},
			},
		},
		{
			name:     "Success. Empty list",
			want:     "",
			infoList: []*models.ImgInfo{},
		},
	}
	for _, tt := range tests {
		got := lofToStr(tt.infoList)
		require.Equal(t, tt.want, got)
	}
}

func Test_imgInfoToStr(t *testing.T) {
	tests := []struct {
		name     string
		infoItem *models.ImgInfo
		want     string
	}{
		{
			name: "Success",
			want: "random.jpeg | 2006-01-02 15:04:05 | 2006-01-02 15:04:05",
			infoItem: &models.ImgInfo{
				FileName: "random.jpeg",
				CreateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
				UpdateDt: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local),
			},
		},
		{
			name:     "Empty",
			want:     "",
			infoItem: &models.ImgInfo{},
		},
	}
	for _, tt := range tests {
		got := imgInfoToStr(tt.infoItem)
		require.Equal(t, tt.want, got)
	}
}
