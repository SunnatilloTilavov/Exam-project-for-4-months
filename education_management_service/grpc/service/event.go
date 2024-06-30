package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"
	"fmt"
	"time"

	"github.com/saidamir98/udevs_pkg/logger"
)

type EventService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedEventServiceServer
}

func NewEventService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventService {
	return &EventService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *EventService) Create(ctx context.Context, req *education_management_service.CreateEventRequest) (*education_management_service.EventResponse, error) {
	s.log.Info("---CreateEvent--->>>", logger.Any("req", req))

	eventDate := req.Date
	t, err := time.Parse("2006-01-02", eventDate)
	if err != nil {
		s.log.Error("---CreateEvent--->>> Invalid date format", logger.Error(err))
		return &education_management_service.EventResponse{}, err
	}

	// Shanba yoki Yakshanba kuniga tekshirish
	if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
		err := fmt.Errorf("Event date must be on a Saturday or Sunday")
		s.log.Error("---CreateEvent--->>>", logger.Error(err))
		return &education_management_service.EventResponse{}, err
	}

	// Event yaratish
	resp, err := s.strg.Event().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateEvent--->>>", logger.Error(err))
		return &education_management_service.EventResponse{}, err
	}

	return resp, nil
}

func (s *EventService) GetByID(ctx context.Context, req *education_management_service.EventID) (*education_management_service.GetEventResponse, error) {
	s.log.Info("---GetEventByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Event().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetEventByID--->>>", logger.Error(err))
		return &education_management_service.GetEventResponse{}, err
	}

	return resp, nil
}

func (s *EventService) GetList(ctx context.Context, req *education_management_service.GetListEventRequest) (*education_management_service.GetListEventResponse, error) {
	s.log.Info("---GetEventList--->>>", logger.Any("req", req))

	resp, err := s.strg.Event().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetEventList--->>>", logger.Error(err))
		return &education_management_service.GetListEventResponse{}, err
	}

	return resp, nil
}

func (s *EventService) Update(ctx context.Context, req *education_management_service.UpdateEventRequest) (*education_management_service.GetEventResponse, error) {
	s.log.Info("---UpdateEvent--->>>", logger.Any("req", req))
	resp, err := s.strg.Event().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateEvent--->>>", logger.Error(err))
		return &education_management_service.GetEventResponse{}, err
	}

	return resp, nil
}

func (s *EventService) Delete(ctx context.Context, req *education_management_service.EventID) (*education_management_service.EventEmpty, error) {
	s.log.Info("---DeleteEvent--->>>", logger.Any("req", req))

	resp, err := s.strg.Event().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteEvent--->>>", logger.Error(err))
		return &education_management_service.EventEmpty{}, err
	}

	return resp, nil
}
