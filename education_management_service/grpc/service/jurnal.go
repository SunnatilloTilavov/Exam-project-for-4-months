package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type JurnalService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedJurnalServiceServer
}

func NewJurnalService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *JurnalService {
	return &JurnalService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *JurnalService) Create(ctx context.Context, req *education_management_service.CreateJurnalRequest) (*education_management_service.JurnalResponse, error) {
	s.log.Info("---CreateJurnal--->>>", logger.Any("req", req))
	resp, err := s.strg.Jurnal().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateJurnal--->>>", logger.Error(err))
		return &education_management_service.JurnalResponse{}, err
	}

	return resp, nil
}

func (s *JurnalService) GetByID(ctx context.Context, req *education_management_service.JurnalID) (*education_management_service.GetJurnalResponse, error) {
	s.log.Info("---GetJurnalByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Jurnal().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetJurnalByID--->>>", logger.Error(err))
		return &education_management_service.GetJurnalResponse{}, err
	}

	return resp, nil
}

func (s *JurnalService) GetList(ctx context.Context, req *education_management_service.GetListJurnalRequest) (*education_management_service.GetListJurnalResponse, error) {
	s.log.Info("---GetJurnalList--->>>", logger.Any("req", req))

	resp, err := s.strg.Jurnal().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetJurnalList--->>>", logger.Error(err))
		return &education_management_service.GetListJurnalResponse{}, err
	}

	return resp, nil
}

func (s *JurnalService) Update(ctx context.Context, req *education_management_service.UpdateJurnalRequest) (*education_management_service.GetJurnalResponse, error) {
	s.log.Info("---UpdateJurnal--->>>", logger.Any("req", req))
	resp, err := s.strg.Jurnal().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateJurnal--->>>", logger.Error(err))
		return &education_management_service.GetJurnalResponse{}, err
	}

	return resp, nil
}

func (s *JurnalService) Delete(ctx context.Context, req *education_management_service.JurnalID) (*education_management_service.JurnalEmpty, error) {
	s.log.Info("---DeleteJurnal--->>>", logger.Any("req", req))

	resp, err := s.strg.Jurnal().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteJurnal--->>>", logger.Error(err))
		return &education_management_service.JurnalEmpty{}, err
	}

	return resp, nil
}


func (s *JurnalService) GetByIDStudent(ctx context.Context, req *education_management_service.GroupId) (*education_management_service.GetJurnalResponse, error) {
	s.log.Info("---GetJurnalByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Jurnal().GetByGroupID(ctx, req)
	if err != nil {
		s.log.Error("---GetJurnalByID--->>>", logger.Error(err))
		return &education_management_service.GetJurnalResponse{}, err
	}

	return resp, nil
}