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

type TeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedTeacherServiceServer
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *TeacherService) Create(ctx context.Context, req *user_service.CreateTeacherRequest) (*user_service.TeacherResponse, error) {
	s.log.Info("---CreateTeacher--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---CreateTeacher--->>>", logger.Error(err))
		return &user_service.TeacherResponse{}, err
	}
	resp, err := s.strg.Teacher().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateTeacher--->>>", logger.Error(err))
		return &user_service.TeacherResponse{}, err
	}

	return resp, nil
}

func (s *TeacherService) GetByID(ctx context.Context, req *user_service.TeacherID) (*user_service.GetTeacherResponse, error) {
	s.log.Info("---GetTeacherByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Teacher().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetTeacherByID--->>>", logger.Error(err))
		return &user_service.GetTeacherResponse{}, err
	}

	return resp, nil
}

func (s *TeacherService) GetList(ctx context.Context, req *user_service.GetListTeacherRequest) (*user_service.GetListTeacherResponse, error) {
	s.log.Info("---GetTeacherList--->>>", logger.Any("req", req))

	resp, err := s.strg.Teacher().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetTeacherList--->>>", logger.Error(err))
		return &user_service.GetListTeacherResponse{}, err
	}

	return resp, nil
}

func (s *TeacherService) Update(ctx context.Context, req *user_service.UpdateTeacherRequest) (*user_service.GetTeacherResponse, error) {
	s.log.Info("---UpdateTeacher--->>>", logger.Any("req", req))

	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---UpdateSupportTeacher--->>>", logger.Error(err))
		return  &user_service.GetTeacherResponse{}, err
	}
	resp, err := s.strg.Teacher().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateTeacher--->>>", logger.Error(err))
		return &user_service.GetTeacherResponse{}, err
	}

	return resp, nil
}

func (s *TeacherService) Delete(ctx context.Context, req *user_service.TeacherID) (*user_service.TeacherEmpty, error) {
	s.log.Info("---DeleteTeacher--->>>", logger.Any("req", req))

	resp, err := s.strg.Teacher().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteTeacher--->>>", logger.Error(err))
		return &user_service.TeacherEmpty{}, err
	}

	return resp, nil
}


func (s *TeacherService) GetReportList(ctx context.Context, req *user_service.GetReportListTeacherRequest) (*user_service.GetReportListTeacherResponse, error) {
	s.log.Info("---GetTeacherList--->>>", logger.Any("req", req))

	resp, err := s.strg.Teacher().GetReportList(ctx, req)
	if err != nil {
		s.log.Error("---GetTeacherList--->>>", logger.Error(err))
		return &user_service.GetReportListTeacherResponse{}, err
	}

	return resp, nil
}