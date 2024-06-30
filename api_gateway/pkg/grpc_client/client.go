package grpc_client

import (
	bc "api_gateway/genproto/education_management_service"
	pc "api_gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	TeacherService() pc.TeacherServiceClient
	SupportTeacherService() pc.SupportTeacherServiceClient
	StudentService() pc.StudentServiceClient
	ManagerService() pc.ManagerServiceClient
	AdministrationService() pc.AdministrationServiceClient
	LoginService() pc.LoginServiceClient
	BranchService() bc.BranchServiceClient
	EventStudentService() bc.EventStudentServiceClient
	EventService() bc.EventServiceClient
	GroupService() bc.GroupServiceClient
	JurnalService() bc.JurnalServiceClient
	ScheduleService() bc.ScheduleServiceClient
	StudentPaymentService() bc.StudentPaymentServiceClient
	StudentTaskService() bc.StudentTaskServiceClient
	TaskService() bc.TaskServiceClient
}

// GrpcClient ...Group
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("User service dial host: %s port:%s err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connMana, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.ManaServiceHost, cfg.ManaServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("User service dial host: %s port:%s err: %s",
			cfg.ManaServiceHost, cfg.ManaServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"Teacher_service":        pc.NewTeacherServiceClient(connUser),
			"SupportTeacher_service": pc.NewSupportTeacherServiceClient(connUser),
			"Student_service":        pc.NewStudentServiceClient(connUser),
			"Manager_service":        pc.NewManagerServiceClient(connUser),
			"Administration_service": pc.NewAdministrationServiceClient(connUser),
			"Login_service":          pc.NewLoginServiceClient(connUser),
			"Branch_service":         bc.NewBranchServiceClient(connMana),
			"EventStudent_service":   bc.NewEventStudentServiceClient(connMana),
			"Event_service":          bc.NewEventServiceClient(connMana),
			"Group_service":          bc.NewGroupServiceClient(connMana),
			"StudentPayment_service": bc.NewStudentPaymentServiceClient(connMana),
			"Jurnal_service":         bc.NewJurnalServiceClient(connMana),
			"Schedule_service":       bc.NewScheduleServiceClient(connMana),
			"StudentTask_service":    bc.NewStudentTaskServiceClient(connMana),
			"Task_service":           bc.NewTaskServiceClient(connMana),
		},
	}, nil
}

func (g *GrpcClient) TeacherService() pc.TeacherServiceClient {
	return g.connections["Teacher_service"].(pc.TeacherServiceClient)
}

func (g *GrpcClient) SupportTeacherService() pc.SupportTeacherServiceClient {
	return g.connections["SupportTeacher_service"].(pc.SupportTeacherServiceClient)
}

func (g *GrpcClient) StudentService() pc.StudentServiceClient {
	return g.connections["Student_service"].(pc.StudentServiceClient)
}

func (g *GrpcClient) ManagerService() pc.ManagerServiceClient {
	return g.connections["Manager_service"].(pc.ManagerServiceClient)
}

func (g *GrpcClient) AdministrationService() pc.AdministrationServiceClient {
	return g.connections["Administration_service"].(pc.AdministrationServiceClient)
}
func (g *GrpcClient) LoginService() pc.LoginServiceClient {
	return g.connections["Login_service"].(pc.LoginServiceClient)
}

func (g *GrpcClient) BranchService() bc.BranchServiceClient {
	return g.connections["Branch_service"].(bc.BranchServiceClient)
}

func (g *GrpcClient) EventStudentService() bc.EventStudentServiceClient {
	return g.connections["EventStudent_service"].(bc.EventStudentServiceClient)
}
func (g *GrpcClient) EventService() bc.EventServiceClient {
	return g.connections["Event_service"].(bc.EventServiceClient)
}

func (g *GrpcClient) GroupService() bc.GroupServiceClient {
	return g.connections["Group_service"].(bc.GroupServiceClient)
}

func (g *GrpcClient) JurnalService() bc.JurnalServiceClient {
	return g.connections["Jurnal_service"].(bc.JurnalServiceClient)
}

func (g *GrpcClient) ScheduleService() bc.ScheduleServiceClient {
	return g.connections["Schedule_service"].(bc.ScheduleServiceClient)
}

func (g *GrpcClient) StudentPaymentService() bc.StudentPaymentServiceClient {
	return g.connections["StudentPayment_service"].(bc.StudentPaymentServiceClient)
}

func (g *GrpcClient) StudentTaskService() bc.StudentTaskServiceClient {
	return g.connections["StudentTask_service"].(bc.StudentTaskServiceClient)
}

func (g *GrpcClient) TaskService() bc.TaskServiceClient {
	return g.connections["Task_service"].(bc.TaskServiceClient)
}