package main

import (
	"context"
	"education_management_service/config"
	"education_management_service/grpc"
	"education_management_service/grpc/client"
	"education_management_service/storage/postgres"
	"net"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log := logger.NewLogger("recoveryLogger", logger.LevelDebug)
			log.Error("Panic recovered", logger.Any("error", r))
		}
	}()

	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	grpcServer := grpc.SetUpServer(cfg, log, pgStore, svcs)

	lis, err := net.Listen("tcp", cfg.ContentGRPCPort)
	if err != nil {
		log.Panic("net.Listen", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.ContentGRPCPort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve", logger.Error(err))
	}
}
