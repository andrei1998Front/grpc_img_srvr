package app

import (
	"log/slog"

	grpcapp "github.com/andrei1998Front/grpc_img_srvr/internal/app/grpc"
	imgService "github.com/andrei1998Front/grpc_img_srvr/internal/services"
	"github.com/andrei1998Front/grpc_img_srvr/internal/storage"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	maxDownloadUploadCalls int,
	maxLOFCalls int,
) *App {
	storage := storage.New(storagePath)

	imgService := imgService.New(log, storage, storage, storage)

	grpcApp := grpcapp.New(log, *imgService, grpcPort, maxDownloadUploadCalls, maxLOFCalls)

	return &App{
		GRPCServer: grpcApp,
	}
}
