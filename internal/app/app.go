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
) *App {
	storage := storage.New(storagePath)

	imgService := imgService.New(log, storage)

	grpcApp := grpcapp.New(log, *imgService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

