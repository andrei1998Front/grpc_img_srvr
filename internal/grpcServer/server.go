package grpcServer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
	gis "github.com/andrei1998Front/grpc_img_srvr/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type serverApi struct {
	gis.UnimplementedImgServiceServer
	imgService       ImgService
	chDownloadUpload chan struct{}
	chLOF            chan struct{}
}

func New(
	imgService ImgService,
	chDownloadUpload chan struct{},
	chLOF chan struct{},
) *serverApi {
	return &serverApi{
		imgService:       imgService,
		chDownloadUpload: chDownloadUpload,
		chLOF:            chLOF,
	}
}

type ImgService interface {
	Upload(
		filename string,
		chunk bytes.Buffer,
	) (models.ImgInfo, error)
	Download(
		filename string,
	) (*models.ImgInfo, error)
	ListOfImages(ctx context.Context) ([]*models.ImgInfo, error)
}

func Register(
	gRPCServer *grpc.Server,
	imgService ImgService,
	chDownloadUpload chan struct{},
	chLOF chan struct{},
) {
	sApi := New(imgService, chDownloadUpload, chLOF)
	gis.RegisterImgServiceServer(gRPCServer, sApi)
}

func (s *serverApi) Upload(
	stream gis.ImgService_UploadServer,
) error {
	s.chDownloadUpload <- struct{}{}
	defer func() { <-s.chDownloadUpload }()

	var filename string
	imageData := bytes.Buffer{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, "internal error")
		}

		filename = req.GetFilename()

		if filename == "" {
			return status.Error(codes.InvalidArgument, "filename is required")
		}

		_, err = imageData.Write(req.GetChunk())
		if err != nil {
			return status.Error(codes.Internal, "internal error")
		}
	}

	fl, err := s.imgService.Upload(filename, imageData)

	if err != nil {
		return status.Error(codes.Internal, "internal error")
	}

	imgInfo := gis.ImgInfo{
		FileName: fl.FileName,
		Size:     fl.Size,
		CreateDt: timestamppb.New(fl.CreateDt),
		UpdateDt: timestamppb.New(fl.UpdateDt),
	}

	if err = stream.SendAndClose(&imgInfo); err != nil {
		return status.Error(codes.Internal, "internal error")
	}

	return nil
}

func (s *serverApi) Download(
	req *gis.ImgDownloadRequest,
	stream gis.ImgService_DownloadServer,
) error {
	s.chDownloadUpload <- struct{}{}
	defer func() { <-s.chDownloadUpload }()

	filename := req.GetFileName()

	if filename == "" {
		return status.Error(codes.InvalidArgument, "filename is required")
	}

	imgInfo, err := s.imgService.Download(filename)

	if err != nil {
		return status.Error(codes.Internal, "internal error")
	}

	imgFile, err := os.Open(imgInfo.Path + "/" + imgInfo.FileName)
	fmt.Println(imgInfo.Path + "/" + imgInfo.FileName)

	if err != nil {
		return status.Error(codes.Internal, "open file error")
	}

	defer imgFile.Close()

	var totalBytesStreamed uint32

	for totalBytesStreamed < imgInfo.Size {
		chunk := make([]byte, 1024)
		bytesRead, err := imgFile.Read(chunk)

		if err == io.EOF {
			break
		}

		if err != nil {
			return status.Error(codes.Internal, "downlod file error")
		}

		if err := stream.Send(&gis.ImgDownloadResponce{
			FileName: imgInfo.FileName,
			Chunk:    chunk,
		}); err != nil {
			return status.Error(codes.Internal, "downlod file error")
		}

		totalBytesStreamed += uint32(bytesRead)
	}

	return nil
}

func (s *serverApi) ListOfImages(ctx context.Context, req *gis.ListOfImagesRequest) (*gis.ListOfImagesResponce, error) {
	s.chLOF <- struct{}{}
	defer func() { <-s.chLOF }()

	lof, err := s.imgService.ListOfImages(ctx)

	if err != nil {
		return &gis.ListOfImagesResponce{}, status.Error(codes.Internal, "internal error")
	}

	var lofResponce gis.ListOfImagesResponce
	for _, imgInfo := range lof {
		lofResponce.ListOfImages = append(lofResponce.ListOfImages, &gis.ImgInfo{
			FileName: imgInfo.FileName,
			Size:     imgInfo.Size,
			CreateDt: timestamppb.New(imgInfo.CreateDt),
			UpdateDt: timestamppb.New(imgInfo.UpdateDt),
		})
	}

	return &lofResponce, nil
}
