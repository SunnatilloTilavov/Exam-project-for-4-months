package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"
	"time"
	"errors"

	"github.com/saidamir98/udevs_pkg/logger"
)

type EventStudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedEventStudentServiceServer
}

func NewEventStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventStudentService {
	return &EventStudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

// func (s *EventStudentService) Create(ctx context.Context, req *education_management_service.CreateEventStudentRequest) (*education_management_service.EventStudentResponse, error) {
// 	s.log.Info("---CreateEventStudent--->>>", logger.Any("req", req))
// 	resp, err := s.strg.EventStudent().Create(ctx, req)
// 	if err != nil {
// 		s.log.Error("---CreateEventStudent--->>>", logger.Error(err))
// 		return &education_management_service.EventStudentResponse{}, err
// 	}

// 	return resp, nil
// }

// func (s *EventStudentService) Create(ctx context.Context, req *education_management_service.CreateEventStudentRequest) (*education_management_service.EventStudentResponse, error) {
// 	s.log.Info("---CreateEventStudent--->>>", logger.Any("req", req))

// 	// Get event by ID
// 	eventResp, err := s.strg.Event().GetByID(ctx, &education_management_service.EventID{Id: req.EventId})
// 	if err != nil {
// 		s.log.Error("---GetEventByID--->>>", logger.Error(err))
// 		return &education_management_service.EventStudentResponse{}, err
// 	}

// 	// Parse event start time
// 	eventStartTime, err := time.Parse("15:04:05", eventResp.StartTime) // Assuming StartTime is in "HH:MM:SS" format
// 	if err != nil {
// 		s.log.Error("---ParseEventStartTime--->>>", logger.Error(err))
// 		return &education_management_service.EventStudentResponse{}, err
// 	}
	
// 	currentTime := time.Now()
// 	eventStartTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), eventStartTime.Hour(), eventStartTime.Minute(), eventStartTime.Second(), 0, currentTime.Location())
// 	if currentTime.After(eventStartTime.Add(-3 * time.Hour)) {
// 		s.log.Error("---EventRegistrationClosed--->>>", logger.Error(errors.New("Event registration closed within 3 hours of event start time")))
// 		return &education_management_service.EventStudentResponse{}, errors.New("Event registration closed within 3 hours of event start time")
// 	}

// 	// Create event student record
// 	resp, err := s.strg.EventStudent().Create(ctx, req)
// 	if err != nil {
// 		s.log.Error("---CreateEventStudent--->>>", logger.Error(err))
// 		return &education_management_service.EventStudentResponse{}, err
// 	}

// 	return resp, nil
// }






func (s *EventStudentService) Create(ctx context.Context, req *education_management_service.CreateEventStudentRequest) (*education_management_service.EventStudentResponse, error) {
	s.log.Info("---CreateEventStudent--->>>", logger.Any("req", req))

	// Get event by ID
	eventResp, err := s.strg.Event().GetByID(ctx, &education_management_service.EventID{Id: req.EventId})
	if err != nil {
		s.log.Error("---GetEventByID--->>>", logger.Error(err))
		return &education_management_service.EventStudentResponse{}, err
	}

	// Parse event start date and time
	startdate:=eventResp.Date[:10]
	eventStartDate, err := time.Parse("2006-01-02", startdate) // Assuming StartDate is in "YYYY-MM-DD" format
	if err != nil {
		s.log.Error("---ParseEventStartDate--->>>", logger.Error(err))
		return &education_management_service.EventStudentResponse{}, err
	}

	eventStartTime, err := time.Parse("15:04:05", eventResp.StartTime) // Assuming StartTime is in "HH:MM:SS" format
	if err != nil {
		s.log.Error("---ParseEventStartTime--->>>", logger.Error(err))
		return &education_management_service.EventStudentResponse{}, err
	}

	// Combine event start date and time
	eventStartDateTime := time.Date(eventStartDate.Year(), eventStartDate.Month(), eventStartDate.Day(), eventStartTime.Hour(), eventStartTime.Minute(), eventStartTime.Second(), 0, time.Local)

	// Current time
	currentTime := time.Now()

	// Check if event is within 3 hours
	registrationCloseTime := eventStartDateTime.Add(-3 * time.Hour)
	if currentTime.After(registrationCloseTime) {
		s.log.Error("---EventRegistrationClosed--->>>", logger.Error(errors.New("Event registration closed within 3 hours of event start time")))
		return &education_management_service.EventStudentResponse{}, errors.New("Event registration closed within 3 hours of event start time")
	}

	// Check if event is today and within 3 hours or in the past
	if eventStartDate.Equal(currentTime) && currentTime.After(registrationCloseTime) {
		s.log.Error("---EventRegistrationClosed--->>>", logger.Error(errors.New("Event registration closed for today's event within 3 hours of event start time")))
		return &education_management_service.EventStudentResponse{}, errors.New("Event registration closed for today's event within 3 hours of event start time")
	}

	resp, err := s.strg.EventStudent().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateEventStudent--->>>", logger.Error(err))
		return &education_management_service.EventStudentResponse{}, err
	}
	
	return resp, nil
}



func (s *EventStudentService) GetByID(ctx context.Context, req *education_management_service.EventStudentID) (*education_management_service.GetEventStudentResponse, error) {
	s.log.Info("---GetEventStudentByID--->>>", logger.Any("req", req))

	resp, err := s.strg.EventStudent().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetEventStudentByID--->>>", logger.Error(err))
		return &education_management_service.GetEventStudentResponse{}, err
	}

	return resp, nil
}

func (s *EventStudentService) GetList(ctx context.Context, req *education_management_service.GetListEventStudentRequest) (*education_management_service.GetListEventStudentResponse, error) {
	s.log.Info("---GetEventStudentList--->>>", logger.Any("req", req))

	resp, err := s.strg.EventStudent().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetEventStudentList--->>>", logger.Error(err))
		return &education_management_service.GetListEventStudentResponse{}, err
	}

	return resp, nil
}

func (s *EventStudentService) Update(ctx context.Context, req *education_management_service.UpdateEventStudentRequest) (*education_management_service.GetEventStudentResponse, error) {
	s.log.Info("---UpdateEventStudent--->>>", logger.Any("req", req))
	resp, err := s.strg.EventStudent().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateEventStudent--->>>", logger.Error(err))
		return &education_management_service.GetEventStudentResponse{}, err
	}

	return resp, nil
}

func (s *EventStudentService) Delete(ctx context.Context, req *education_management_service.EventStudentID) (*education_management_service.EventStudentEmpty, error) {
	s.log.Info("---DeleteEventStudent--->>>", logger.Any("req", req))

	resp, err := s.strg.EventStudent().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteEventStudent--->>>", logger.Error(err))
		return &education_management_service.EventStudentEmpty{}, err
	}

	return resp, nil
}

func (s *EventStudentService) GetStudentByID(ctx context.Context, req *education_management_service.StudentID) (*education_management_service.GetStudentWithEventsResponse, error) {
	s.log.Info("---GetStudentByID--->>>", logger.Any("req", req))

	resp, err := s.strg.EventStudent().GetStudentWithEventsByID(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentByID--->>>", logger.Error(err))
		return &education_management_service.GetStudentWithEventsResponse{}, err
	}
	return resp, nil
}