package tests

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	gis "github.com/andrei1998Front/grpc_img_srvr/pkg/proto"
	"github.com/andrei1998Front/grpc_img_srvr/tests/suite"
	"github.com/stretchr/testify/require"
)

func TestListOfImages(t *testing.T) {
	ctx, st := suite.New(t)

	_, err := st.ImgServiceClient.ListOfImages(ctx, &gis.ListOfImagesRequest{})
	require.NoError(t, err)
}

func TestDownload(t *testing.T) {
	_, st := suite.New(t)

	tests := []struct {
		name      string
		wantError bool
		filename  string
		err       string
	}{
		{
			name:      "Non-existing file",
			wantError: true,
			filename:  "dfff.gg",
			err:       "internal error",
		},
		{
			name:      "empty filename",
			wantError: true,
			filename:  "",
			err:       "filename is required",
		},
		{
			name:      "success",
			wantError: true,
			filename:  "cat.jpg",
			err:       "filename is required",
		},
	}

	for _, tt := range tests {
		err := download(tt.filename, st.ImgServiceClient)

		if tt.wantError && err != nil {
			require.ErrorContains(t, err, tt.err)
		} else if !tt.wantError && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		}
	}
}

func download(fileName string, client gis.ImgServiceClient) error {
	req := &gis.ImgDownloadRequest{
		FileName: fileName,
	}

	stream, err := client.Download(context.Background(), req)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			if err := os.WriteFile(fileName, buffer.Bytes(), 0777); err != nil {
				return err
			}
			break
		}
		if err != nil {
			buffer.Reset()
			return err
		}

		chunk := res.GetChunk()

		buffer.Write(chunk)
	}
	return nil
}

func TestUpload(t *testing.T) {
	_, st := suite.New(t)

	tests := []struct {
		name      string
		filename  string
		wantError bool
		err       string
	}{
		{
			name:      "Success",
			wantError: false,
			filename:  "test.jpg",
		},
		{
			name:      "empty filename",
			wantError: true,
			filename:  "",
			err:       "internal error",
		},
		{
			name:      "existing image",
			wantError: true,
			filename:  "test.jpg",
			err:       "internal error",
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)
		err := upload(tt.filename, st.ImgServiceClient)

		if tt.wantError && err != nil {
			require.ErrorContains(t, err, tt.err)
		} else if !tt.wantError && err != nil {
			t.Errorf("unexpected error: " + err.Error())
			return
		}

		_, err = os.Stat("../imgs/" + tt.filename)

		if err != nil {
			t.Errorf("Image not uploaded: " + err.Error())
			return
		}
	}
}

func upload(filename string, client gis.ImgServiceClient) error {
	stream, err := client.Upload(context.Background())

	if err != nil {
		return err
	}

	imgStat, err := os.Stat("./" + filename)

	if err != nil {
		return err
	}

	img, err := os.Open("./" + filename)

	if err != nil {
		return err
	}

	defer img.Close()

	var totalBytesStreamed int64

	for totalBytesStreamed < imgStat.Size() {
		chunk := make([]byte, 1024)
		bytesRead, err := img.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err := stream.Send(&gis.ImgUploadRequest{
			Filename: filename,
			Chunk:    chunk,
		}); err != nil {
			return err
		}

		totalBytesStreamed += int64(bytesRead)
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}
