package grpcServer

import (
	"bytes"
	"io"

	"github.com/andrei1998Front/grpc_img_srvr/internal/domain/models"
	gis "github.com/andrei1998Front/grpc_img_srvr/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type serverApi struct {
	gis.UnimplementedImgServiceServer
	imgService ImgService
}

type ImgService interface {
	Upload(
		filename string,
		chunk bytes.Buffer,
	) (models.ImgInfo, error)
}

func Register(gRPCServer *grpc.Server, imgService ImgService) {
	gis.RegisterImgServiceServer(gRPCServer, &serverApi{imgService: imgService})
}

func (s *serverApi) Upload(
	stream gis.ImgService_UploadServer,
) error {
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
			return status.Error(codes.Internal, "filename is required")
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
