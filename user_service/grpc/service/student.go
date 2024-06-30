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

type StudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*user_service.UnimplementedStudentServiceServer
}

func NewStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *StudentService) Create(ctx context.Context, req *user_service.CreateStudentRequest) (*user_service.StudentResponse, error) {
	s.log.Info("---CreateStudent--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---CreateStudent--->>>", logger.Error(err))
		return &user_service.StudentResponse{}, err
	}
	resp, err := s.strg.Student().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateStudent--->>>", logger.Error(err))
		return &user_service.StudentResponse{}, err
	}

	return resp, nil
}

func (s *StudentService) GetByID(ctx context.Context, req *user_service.StudentID) (*user_service.GetStudentResponse, error) {
	s.log.Info("---GetStudentByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Student().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentByID--->>>", logger.Error(err))
		return &user_service.GetStudentResponse{}, err
	}

	return resp, nil
}

func (s *StudentService) GetList(ctx context.Context, req *user_service.GetListStudentRequest) (*user_service.GetListStudentResponse, error) {
	s.log.Info("---GetStudentList--->>>", logger.Any("req", req))

	resp, err := s.strg.Student().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentList--->>>", logger.Error(err))
		return &user_service.GetListStudentResponse{}, err
	}

	return resp, nil
}

func (s *StudentService) Update(ctx context.Context, req *user_service.UpdateStudentRequest) (*user_service.GetStudentResponse, error) {
	s.log.Info("---UpdateStudent--->>>", logger.Any("req", req))
	password, err := password.HashPassword(req.Password)
	req.Password=password
	if err != nil {
		s.log.Error("---UpdateStudent--->>>", logger.Error(err))
		return  &user_service.GetStudentResponse{}, err
	}
	resp, err := s.strg.Student().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudent--->>>", logger.Error(err))
		return &user_service.GetStudentResponse{}, err
	}

	return resp, nil
}

func (s *StudentService) Delete(ctx context.Context, req *user_service.StudentID) (*user_service.StudentEmpty, error) {
	s.log.Info("---DeleteStudent--->>>", logger.Any("req", req))

	resp, err := s.strg.Student().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteStudent--->>>", logger.Error(err))
		return &user_service.StudentEmpty{}, err
	}

	return resp, nil
}

func (s *StudentService) GetReportList(ctx context.Context, req *user_service.GetReportListStudentRequest) (*user_service.GetReportListStudentResponse, error) {
	s.log.Info("---GetStudentList--->>>", logger.Any("req", req))

	resp, err := s.strg.Student().GetReportList(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentList--->>>", logger.Error(err))
		return &user_service.GetReportListStudentResponse{}, err
	}

	return resp, nil
}

