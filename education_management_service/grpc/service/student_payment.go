package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type StudentPaymentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedStudentPaymentServiceServer
}

func NewStudentPaymentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentPaymentService {
	return &StudentPaymentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *StudentPaymentService) Create(ctx context.Context, req *education_management_service.CreateStudentPaymentRequest) (*education_management_service.StudentPaymentResponse, error) {
	s.log.Info("---CreateStudentPayment--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentPayment().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateStudentPayment--->>>", logger.Error(err))
		return &education_management_service.StudentPaymentResponse{}, err
	}

	return resp, nil
}

func (s *StudentPaymentService) GetByID(ctx context.Context, req *education_management_service.StudentPaymentID) (*education_management_service.GetStudentPaymentResponse, error) {
	s.log.Info("---GetStudentPaymentByID--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentPayment().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentPaymentByID--->>>", logger.Error(err))
		return &education_management_service.GetStudentPaymentResponse{}, err
	}

	return resp, nil
}

func (s *StudentPaymentService) GetList(ctx context.Context, req *education_management_service.GetListStudentPaymentRequest) (*education_management_service.GetListStudentPaymentResponse, error) {
	s.log.Info("---GetStudentPaymentList--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentPayment().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetStudentPaymentList--->>>", logger.Error(err))
		return &education_management_service.GetListStudentPaymentResponse{}, err
	}

	return resp, nil
}

func (s *StudentPaymentService) Update(ctx context.Context, req *education_management_service.UpdateStudentPaymentRequest) (*education_management_service.GetStudentPaymentResponse, error) {
	s.log.Info("---UpdateStudentPayment--->>>", logger.Any("req", req))
	resp, err := s.strg.StudentPayment().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentPayment--->>>", logger.Error(err))
		return &education_management_service.GetStudentPaymentResponse{}, err
	}

	return resp, nil
}

func (s *StudentPaymentService) Delete(ctx context.Context, req *education_management_service.StudentPaymentID) (*education_management_service.StudentPaymentEmpty, error) {
	s.log.Info("---DeleteStudentPayment--->>>", logger.Any("req", req))

	resp, err := s.strg.StudentPayment().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteStudentPayment--->>>", logger.Error(err))
		return &education_management_service.StudentPaymentEmpty{}, err
	}

	return resp, nil
}
