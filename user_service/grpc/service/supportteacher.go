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

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedSupportTeacherServiceServer
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *SupportTeacherService) Create(ctx context.Context, req *user_service.CreateSupportTeacherRequest) (*user_service.SupportTeacherResponse, error) {
	s.log.Info("---CreateSupportTeacher--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---CreateSupportTeacher--->>>", logger.Error(err))
		return &user_service.SupportTeacherResponse{}, err
	}

	resp, err := s.strg.SupportTeacher().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateSupportTeacher--->>>", logger.Error(err))
		return &user_service.SupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SupportTeacherService) GetByID(ctx context.Context, req *user_service.SupportTeacherID) (*user_service.GetSupportTeacherResponse, error) {
	s.log.Info("---GetSupportTeacherByID--->>>", logger.Any("req", req))

	resp, err := s.strg.SupportTeacher().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetSupportTeacherByID--->>>", logger.Error(err))
		return &user_service.GetSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SupportTeacherService) GetList(ctx context.Context, req *user_service.GetListSupportTeacherRequest) (*user_service.GetListSupportTeacherResponse, error) {
	s.log.Info("---GetSupportTeacherList--->>>", logger.Any("req", req))

	resp, err := s.strg.SupportTeacher().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetSupportTeacherList--->>>", logger.Error(err))
		return &user_service.GetListSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SupportTeacherService) Update(ctx context.Context, req *user_service.UpdateSupportTeacherRequest) (*user_service.GetSupportTeacherResponse, error) {
	s.log.Info("---UpdateSupportTeacher--->>>", logger.Any("req", req))

	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---UpdateSupportTeacher--->>>", logger.Error(err))
		return  &user_service.GetSupportTeacherResponse{}, err
	}
	resp, err := s.strg.SupportTeacher().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateSupportTeacher--->>>", logger.Error(err))
		return &user_service.GetSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SupportTeacherService) Delete(ctx context.Context, req *user_service.SupportTeacherID) (*user_service.SupportTeacherEmpty, error) {
	s.log.Info("---DeleteSupportTeacher--->>>", logger.Any("req", req))

	resp, err := s.strg.SupportTeacher().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteSupportTeacher--->>>", logger.Error(err))
		return &user_service.SupportTeacherEmpty{}, err
	}

	return resp, nil
}


func (s *SupportTeacherService) GetReportList(ctx context.Context, req *user_service.GetReportListSupportTeacherRequest) (*user_service.GetReportListSupportTeacherResponse, error) {
	s.log.Info("---GetSupportTeacherList--->>>", logger.Any("req", req))

	resp, err := s.strg.SupportTeacher().GetReportList(ctx, req)
	if err != nil {
		s.log.Error("---GetSupportTeacherList--->>>", logger.Error(err))
		return &user_service.GetReportListSupportTeacherResponse{}, err
	}

	return resp, nil
}