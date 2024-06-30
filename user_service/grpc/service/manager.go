package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/pkg/password"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedManagerServiceServer
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *ManagerService) Create(ctx context.Context, req *user_service.CreateManagerRequest) (*user_service.ManagerResponse, error) {
	s.log.Info("---CreateManager--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---CreateManager--->>>", logger.Error(err))
		return &user_service.ManagerResponse{}, err
	}

	resp, err := s.strg.Manager().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateManager--->>>", logger.Error(err))
		return &user_service.ManagerResponse{}, err
	}

	return resp, nil
}

func (s *ManagerService) GetByID(ctx context.Context, req *user_service.ManagerID) (*user_service.GetManagerResponse, error) {
	s.log.Info("---GetManagerByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Manager().GetByID(ctx, req)

	if err != nil {
		s.log.Error("---GetManagerByID--->>>", logger.Error(err))
		return &user_service.GetManagerResponse{}, err
	}

	return resp, nil

}

func (s *ManagerService) GetList(ctx context.Context, req *user_service.GetListManagerRequest) (*user_service.GetListManagerResponse, error) {
	s.log.Info("---GetManagerList--->>>", logger.Any("req", req))

	resp, err := s.strg.Manager().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetManagerList--->>>", logger.Error(err))
		return &user_service.GetListManagerResponse{}, err
	}

	return resp, nil
}

func (s *ManagerService) Update(ctx context.Context, req *user_service.UpdateManagerRequest) (*user_service.GetManagerResponse, error) {
	s.log.Info("---UpdateManager--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---UpdateManager--->>>", logger.Error(err))
		return  &user_service.GetManagerResponse{}, err
	}

	resp, err := s.strg.Manager().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateManager--->>>", logger.Error(err))
		return &user_service.GetManagerResponse{}, err
	}

	return resp, nil
}

func (s *ManagerService) Delete(ctx context.Context, req *user_service.ManagerID) (*user_service.ManagerEmpty, error) {
	s.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	resp, err := s.strg.Manager().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteManager--->>>", logger.Error(err))
		return &user_service.ManagerEmpty{}, err
	}

	return resp, nil
}
