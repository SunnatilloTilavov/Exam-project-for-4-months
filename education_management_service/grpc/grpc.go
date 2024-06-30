package grpc

import (
	"education_management_service/config"
	"education_management_service/genproto/education_management_service"
	"education_management_service/grpc/client"
	"education_management_service/grpc/service"
	"education_management_service/storage"


	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer( cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()
		education_management_service.RegisterBranchServiceServer(grpcServer,service.NewBranchService(cfg,log,strg,srvc))
		education_management_service.RegisterGroupServiceServer(grpcServer,service.NewGroupService(cfg,log,strg,srvc))	
		education_management_service.RegisterJurnalServiceServer(grpcServer,service.NewJurnalService(cfg,log,strg,srvc))
		education_management_service.RegisterScheduleServiceServer(grpcServer,service.NewScheduleService(cfg,log,strg,srvc))	
		education_management_service.RegisterTaskServiceServer(grpcServer,service.NewTaskService(cfg,log,strg,srvc))	
		education_management_service.RegisterStudentTaskServiceServer(grpcServer,service.NewStudentTaskService(cfg,log,strg,srvc))
		education_management_service.RegisterEventServiceServer(grpcServer, service.NewEventService(cfg, log, strg, srvc))	
		education_management_service.RegisterEventStudentServiceServer(grpcServer, service.NewEventStudentService(cfg, log, strg, srvc))
		education_management_service.RegisterStudentPaymentServiceServer(grpcServer,service.NewStudentPaymentService(cfg,log,strg,srvc))
		reflection.Register(grpcServer)
	return
}
