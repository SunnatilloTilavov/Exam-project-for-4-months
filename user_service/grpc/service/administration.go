package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"
	"user_service/pkg/password"

	"github.com/saidamir98/udevs_pkg/logger"
)

type AdministrationService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedAdministrationServiceServer
}

func NewAdministrationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AdministrationService {
	return &AdministrationService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *AdministrationService) Create(ctx context.Context, req *user_service.CreateAdministrationRequest) (*user_service.AdministrationResponse, error) {
	s.log.Info("---CreateAdministration--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---CreateAdministration--->>>", logger.Error(err))
		return &user_service.AdministrationResponse{}, err
	}

	resp, err := s.strg.Administration().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateAdministration--->>>", logger.Error(err))
		return &user_service.AdministrationResponse{}, err
	}

	return resp, nil
}

func (s *AdministrationService) GetByID(ctx context.Context, req *user_service.AdministrationID) (*user_service.GetAdministrationResponse, error) {
	s.log.Info("---GetAdministrationByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Administration().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetAdministrationByID--->>>", logger.Error(err))
		return &user_service.GetAdministrationResponse{}, err
	}

	return resp, nil
}

func (s *AdministrationService) GetList(ctx context.Context, req *user_service.GetListAdministrationRequest) (*user_service.GetListAdministrationResponse, error) {
	s.log.Info("---GetAdministrationList--->>>", logger.Any("req", req))

	resp, err := s.strg.Administration().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetAdministrationList--->>>", logger.Error(err))
		return &user_service.GetListAdministrationResponse{}, err
	}

	return resp, nil
}

func (s *AdministrationService) Update(ctx context.Context, req *user_service.UpdateAdministrationRequest) (*user_service.GetAdministrationResponse, error) {
	s.log.Info("---UpdateAdministration--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---UpdateAdminstration--->>>", logger.Error(err))
		return  &user_service.GetAdministrationResponse{}, err
	}
	resp, err := s.strg.Administration().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateAdministration--->>>", logger.Error(err))
		return &user_service.GetAdministrationResponse{}, err
	}

	return resp, nil
}

func (s *AdministrationService) Delete(ctx context.Context, req *user_service.AdministrationID) (*user_service.AdministrationEmpty, error) {
	s.log.Info("---DeleteAdministration--->>>", logger.Any("req", req))

	resp, err := s.strg.Administration().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteAdministration--->>>", logger.Error(err))
		return &user_service.AdministrationEmpty{}, err
	}

	return resp, nil
}


func (s *AdministrationService) GetReportList(ctx context.Context, req *user_service.GetReportListAdministrationRequest) (*user_service.GetReportListAdministrationResponse, error) {
	s.log.Info("---GetReportAdministrationList--->>>", logger.Any("req", req))

	resp, err := s.strg.Administration().GetReportList(ctx, req)
	if err != nil {
		s.log.Error("---GetReportAdministrationList--->>>", logger.Error(err))
		return &user_service.GetReportListAdministrationResponse{}, err
	}

	return resp, nil
}