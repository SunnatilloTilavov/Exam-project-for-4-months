package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *BranchService) Create(ctx context.Context, req *education_management_service.CreateBranchRequest) (*education_management_service.BranchResponse, error) {
	s.log.Info("---CreateBranch--->>>", logger.Any("req", req))
	resp, err := s.strg.Branch().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateBranch--->>>", logger.Error(err))
		return &education_management_service.BranchResponse{}, err
	}

	return resp, nil
}

func (s *BranchService) GetByID(ctx context.Context, req *education_management_service.BranchID) (*education_management_service.GetBranchResponse, error) {
	s.log.Info("---GetBranchByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Branch().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetBranchByID--->>>", logger.Error(err))
		return &education_management_service.GetBranchResponse{}, err
	}

	return resp, nil
}

func (s *BranchService) GetList(ctx context.Context, req *education_management_service.GetListBranchRequest) (*education_management_service.GetListBranchResponse, error) {
	s.log.Info("---GetBranchList--->>>", logger.Any("req", req))

	resp, err := s.strg.Branch().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetBranchList--->>>", logger.Error(err))
		return &education_management_service.GetListBranchResponse{}, err
	}

	return resp, nil
}

func (s *BranchService) Update(ctx context.Context, req *education_management_service.UpdateBranchRequest) (*education_management_service.GetBranchResponse, error) {
	s.log.Info("---UpdateBranch--->>>", logger.Any("req", req))
	resp, err := s.strg.Branch().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateBranch--->>>", logger.Error(err))
		return &education_management_service.GetBranchResponse{}, err
	}

	return resp, nil
}

func (s *BranchService) Delete(ctx context.Context, req *education_management_service.BranchID) (*education_management_service.BranchEmpty, error) {
	s.log.Info("---DeleteBranch--->>>", logger.Any("req", req))

	resp, err := s.strg.Branch().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteBranch--->>>", logger.Error(err))
		return &education_management_service.BranchEmpty{}, err
	}

	return resp, nil
}
