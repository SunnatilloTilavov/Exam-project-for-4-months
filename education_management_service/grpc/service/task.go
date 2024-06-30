package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type TaskService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedTaskServiceServer
}

func NewTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TaskService {
	return &TaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *TaskService) Create(ctx context.Context, req *education_management_service.CreateTaskRequest) (*education_management_service.TaskResponse, error) {
	s.log.Info("---CreateTask--->>>", logger.Any("req", req))
	resp, err := s.strg.Task().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateTask--->>>", logger.Error(err))
		return &education_management_service.TaskResponse{}, err
	}

	return resp, nil
}

func (s *TaskService) GetByID(ctx context.Context, req *education_management_service.TaskID) (*education_management_service.GetTaskResponse, error) {
	s.log.Info("---GetTaskByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Task().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetTaskByID--->>>", logger.Error(err))
		return &education_management_service.GetTaskResponse{}, err
	}

	return resp, nil
}

func (s *TaskService) GetList(ctx context.Context, req *education_management_service.GetListTaskRequest) (*education_management_service.GetListTaskResponse, error) {
	s.log.Info("---GetTaskList--->>>", logger.Any("req", req))

	resp, err := s.strg.Task().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetTaskList--->>>", logger.Error(err))
		return &education_management_service.GetListTaskResponse{}, err
	}

	return resp, nil
}

func (s *TaskService) Update(ctx context.Context, req *education_management_service.UpdateTaskRequest) (*education_management_service.GetTaskResponse, error) {
	s.log.Info("---UpdateTask--->>>", logger.Any("req", req))
	resp, err := s.strg.Task().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateTask--->>>", logger.Error(err))
		return &education_management_service.GetTaskResponse{}, err
	}

	return resp, nil
}

func (s *TaskService) Delete(ctx context.Context, req *education_management_service.TaskID) (*education_management_service.TaskEmpty, error) {
	s.log.Info("---DeleteTask--->>>", logger.Any("req", req))

	resp, err := s.strg.Task().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteTask--->>>", logger.Error(err))
		return &education_management_service.TaskEmpty{}, err
	}

	return resp, nil
}
