package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ScheduleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedScheduleServiceServer
}

func NewScheduleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScheduleService {
	return &ScheduleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *ScheduleService) Create(ctx context.Context, req *education_management_service.CreateScheduleRequest) (*education_management_service.ScheduleResponse, error) {
	s.log.Info("---CreateSchedule--->>>", logger.Any("req", req))
	resp, err := s.strg.Schedule().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateSchedule--->>>", logger.Error(err))
		return &education_management_service.ScheduleResponse{}, err
	}

	return resp, nil
}

func (s *ScheduleService) GetByID(ctx context.Context, req *education_management_service.ScheduleID) (*education_management_service.GetScheduleResponse, error) {
	s.log.Info("---GetScheduleByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Schedule().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetScheduleByID--->>>", logger.Error(err))
		return &education_management_service.GetScheduleResponse{}, err
	}

	return resp, nil
}

func (s *ScheduleService) GetList(ctx context.Context, req *education_management_service.GetListScheduleRequest) (*education_management_service.GetListScheduleResponse, error) {
	s.log.Info("---GetScheduleList--->>>", logger.Any("req", req))

	resp, err := s.strg.Schedule().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetScheduleList--->>>", logger.Error(err))
		return &education_management_service.GetListScheduleResponse{}, err
	}

	return resp, nil
}

func (s *ScheduleService) Update(ctx context.Context, req *education_management_service.UpdateScheduleRequest) (*education_management_service.GetScheduleResponse, error) {
	s.log.Info("---UpdateSchedule--->>>", logger.Any("req", req))
	resp, err := s.strg.Schedule().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateSchedule--->>>", logger.Error(err))
		return &education_management_service.GetScheduleResponse{}, err
	}

	return resp, nil
}

func (s *ScheduleService) Delete(ctx context.Context, req *education_management_service.ScheduleID) (*education_management_service.ScheduleEmpty, error) {
	s.log.Info("---DeleteSchedule--->>>", logger.Any("req", req))

	resp, err := s.strg.Schedule().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteSchedule--->>>", logger.Error(err))
		return &education_management_service.ScheduleEmpty{}, err
	}

	return resp, nil
}
