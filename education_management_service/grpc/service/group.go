package service

import (
	"context"
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type GroupService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*education_management_service.UnimplementedGroupServiceServer
}

func NewGroupService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupService {
	return &GroupService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *GroupService) Create(ctx context.Context, req *education_management_service.CreateGroupRequest) (*education_management_service.GroupResponse, error) {
	s.log.Info("---CreateGroup--->>>", logger.Any("req", req))
	resp, err := s.strg.Group().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateGroup--->>>", logger.Error(err))
		return &education_management_service.GroupResponse{}, err
	}

	return resp, nil
}

func (s *GroupService) GetByID(ctx context.Context, req *education_management_service.GroupID) (*education_management_service.GetGroupResponse, error) {
	s.log.Info("---GetGroupByID--->>>", logger.Any("req", req))

	resp, err := s.strg.Group().GetByID(ctx, req)
	if err != nil {
		s.log.Error("---GetGroupByID--->>>", logger.Error(err))
		return &education_management_service.GetGroupResponse{}, err
	}

	return resp, nil
}

func (s *GroupService) GetList(ctx context.Context, req *education_management_service.GetListGroupRequest) (*education_management_service.GetListGroupResponse, error) {
	s.log.Info("---GetGroupList--->>>", logger.Any("req", req))

	resp, err := s.strg.Group().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetGroupList--->>>", logger.Error(err))
		return &education_management_service.GetListGroupResponse{}, err
	}

	return resp, nil
}

func (s *GroupService) Update(ctx context.Context, req *education_management_service.UpdateGroupRequest) (*education_management_service.GetGroupResponse, error) {
	s.log.Info("---UpdateGroup--->>>", logger.Any("req", req))
	resp, err := s.strg.Group().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateGroup--->>>", logger.Error(err))
		return &education_management_service.GetGroupResponse{}, err
	}

	return resp, nil
}

func (s *GroupService) Delete(ctx context.Context, req *education_management_service.GroupID) (*education_management_service.GroupEmpty, error) {
	s.log.Info("---DeleteGroup--->>>", logger.Any("req", req))

	resp, err := s.strg.Group().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteGroup--->>>", logger.Error(err))
		return &education_management_service.GroupEmpty{}, err
	}

	return resp, nil
}


func (s *GroupService) GetByIDTeacher(ctx context.Context, req *education_management_service.TeacherID) (*education_management_service.GetGroupResponse, error) {
	s.log.Info("---GetGroupByIDTeacher--->>>", logger.Any("req", req))

	resp, err := s.strg.Group().GetByIDTeacher(ctx, req)
	if err != nil {
		s.log.Error("---GetGroupByIDTeacher--->>>", logger.Error(err))
		return &education_management_service.GetGroupResponse{}, err
	}

	return resp, nil
}