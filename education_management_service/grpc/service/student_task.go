package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type StudentTaskService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedStudentTaskServiceServer
}

func NewStudentTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentTaskService {
	return &StudentTaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *StudentTaskService) Create(ctx context.Context, req *education_management_service.CreateStudentTaskRequest) (*education_management_service.StudentTaskResponse, error) {
	s.log.Info("---CreateStudentTask--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentTask().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateStudentTask--->>>", logger.Error(err))
		return &education_management_service.StudentTaskResponse{}, err
	}

	return resp, nil
}

func (s *StudentTaskService) GetByID(ctx context.Context, req *education_management_service.StudentTaskID) (*education_management_service.GetStudentTaskResponse, error) {
	s.log.Info("---GetStudentTaskByID--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentTask().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentTaskByID--->>>", logger.Error(err))
		return &education_management_service.GetStudentTaskResponse{}, err
	}

	return resp, nil
}

func (s *StudentTaskService) GetList(ctx context.Context, req *education_management_service.GetListStudentTaskRequest) (*education_management_service.GetListStudentTaskResponse, error) {
	s.log.Info("---GetStudentTaskList--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentTask().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentTaskList--->>>", logger.Error(err))
		return &education_management_service.GetListStudentTaskResponse{}, err
	}

	return resp, nil
}

func (s *StudentTaskService) Update(ctx context.Context, req *education_management_service.UpdateStudentTaskRequest) (*education_management_service.GetStudentTaskResponse, error) {
	s.log.Info("---UpdateStudentTask--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentTask().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentTask--->>>", logger.Error(err))
		return &education_management_service.GetStudentTaskResponse{}, err
	}
	return resp, nil
}

func (s *StudentTaskService) Delete(ctx context.Context, req *education_management_service.StudentTaskID) (*education_management_service.StudentTaskEmpty, error) {
	s.log.Info("---DeleteStudentTask--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentTask().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteStudentTask--->>>", logger.Error(err))
		return &education_management_service.StudentTaskEmpty{}, err
	}

	return resp, nil
}

func (s *StudentTaskService) UpdateScoreforTeacher(ctx context.Context, req *education_management_service.UpdateStudentScoreRequest) (*education_management_service.GetStudentTaskResponse, error) {
	s.log.Info("---UpdateStudentTask--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentTask().UpdateScoreforTeacher(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentTask--->>>", logger.Error(err))
		return &education_management_service.GetStudentTaskResponse{}, err
	}
	return resp, nil
}

func (s *StudentTaskService)UpdateScoreforStudent(ctx context.Context, req *education_management_service.UpdateStudentScoreRequest) (*education_management_service.GetStudentTaskResponse, error) {
	s.log.Info("---UpdateStudentTask--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentTask().UpdateScoreforStudent(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentTask--->>>", logger.Error(err))
		return &education_management_service.GetStudentTaskResponse{}, err
	}
	return resp, nil
}

