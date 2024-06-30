package storage

import (
	"context"
	ct "education_management_service/genproto/education_management_service"
)

type StorageI interface {
	CloseDB()

	Branch() BranchRepoI
	Group() GroupRepoI
	Jurnal() JurnalRepoI
	Schedule() ScheduleRepoI
	Task() TaskRepoI
	StudentTask() StudentTaskRepoI
	Event() EventRepoI  
	EventStudent() EventStudentRepoI
	StudentPayment() StudentPaymentRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *ct.CreateBranchRequest) (*ct.BranchResponse, error)
	GetByID(ctx context.Context, req *ct.BranchID) (*ct.GetBranchResponse, error)
	GetList(ctx context.Context, req *ct.GetListBranchRequest) (*ct.GetListBranchResponse, error)
	Update(ctx context.Context, req *ct.UpdateBranchRequest) (*ct.GetBranchResponse, error)
	Delete(ctx context.Context, req *ct.BranchID) (*ct.BranchEmpty, error)
}

type GroupRepoI interface {
	Create(ctx context.Context, req *ct.CreateGroupRequest) (*ct.GroupResponse, error)
	GetByID(ctx context.Context, req *ct.GroupID) (*ct.GetGroupResponse, error)
	GetList(ctx context.Context, req *ct.GetListGroupRequest) (*ct.GetListGroupResponse, error)
	Update(ctx context.Context, req *ct.UpdateGroupRequest) (*ct.GetGroupResponse, error)
	Delete(ctx context.Context, req *ct.GroupID) (*ct.GroupEmpty, error)
	GetByIDTeacher(ctx context.Context, req *ct.TeacherID) (*ct.GetGroupResponse, error)
}

type JurnalRepoI interface {
	Create(ctx context.Context, req *ct.CreateJurnalRequest) (*ct.JurnalResponse, error)
	GetByID(ctx context.Context, req *ct.JurnalID) (*ct.GetJurnalResponse, error)
	GetList(ctx context.Context, req *ct.GetListJurnalRequest) (*ct.GetListJurnalResponse, error)
	Update(ctx context.Context, req *ct.UpdateJurnalRequest) (*ct.GetJurnalResponse, error)
	Delete(ctx context.Context, req *ct.JurnalID) (*ct.JurnalEmpty, error)
	GetByGroupID(ctx context.Context, req *ct.GroupId) (*ct.GetJurnalResponse, error)
}

type ScheduleRepoI interface {
	Create(ctx context.Context, req *ct.CreateScheduleRequest) (*ct.ScheduleResponse, error)
	GetByID(ctx context.Context, req *ct.ScheduleID) (*ct.GetScheduleResponse, error)
	GetList(ctx context.Context, req *ct.GetListScheduleRequest) (*ct.GetListScheduleResponse, error)
	Update(ctx context.Context, req *ct.UpdateScheduleRequest) (*ct.GetScheduleResponse, error)
	Delete(ctx context.Context, req *ct.ScheduleID) (*ct.ScheduleEmpty, error)
}

type TaskRepoI interface {
	Create(ctx context.Context, req *ct.CreateTaskRequest) (*ct.TaskResponse, error)
	GetByID(ctx context.Context, req *ct.TaskID) (*ct.GetTaskResponse, error)
	GetList(ctx context.Context, req *ct.GetListTaskRequest) (*ct.GetListTaskResponse, error)
	Update(ctx context.Context, req *ct.UpdateTaskRequest) (*ct.GetTaskResponse, error)
	Delete(ctx context.Context, req *ct.TaskID) (*ct.TaskEmpty, error)
}

type StudentTaskRepoI interface {
	Create(ctx context.Context, req *ct.CreateStudentTaskRequest) (*ct.StudentTaskResponse, error)
	GetByID(ctx context.Context, req *ct.StudentTaskID) (*ct.GetStudentTaskResponse, error)
	GetList(ctx context.Context, req *ct.GetListStudentTaskRequest) (*ct.GetListStudentTaskResponse, error)
	Update(ctx context.Context, req *ct.UpdateStudentTaskRequest) (*ct.GetStudentTaskResponse, error)
	UpdateScoreforTeacher(ctx context.Context, req *ct.UpdateStudentScoreRequest)(*ct.GetStudentTaskResponse, error)
	UpdateScoreforStudent(ctx context.Context, req *ct.UpdateStudentScoreRequest)(*ct.GetStudentTaskResponse, error)
	Delete(ctx context.Context, req *ct.StudentTaskID) (*ct.StudentTaskEmpty, error)
}

type EventRepoI interface {
	Create(ctx context.Context, req *ct.CreateEventRequest) (*ct.EventResponse, error)
	GetByID(ctx context.Context, req *ct.EventID) (*ct.GetEventResponse, error)
	GetList(ctx context.Context, req *ct.GetListEventRequest) (*ct.GetListEventResponse, error)
	Update(ctx context.Context, req *ct.UpdateEventRequest) (*ct.GetEventResponse, error)
	Delete(ctx context.Context, req *ct.EventID) (*ct.EventEmpty, error)	
}

type EventStudentRepoI interface {
	Create(ctx context.Context, req *ct.CreateEventStudentRequest) (*ct.EventStudentResponse, error)
	GetByID(ctx context.Context, req *ct.EventStudentID) (*ct.GetEventStudentResponse, error)
	GetList(ctx context.Context, req *ct.GetListEventStudentRequest) (*ct.GetListEventStudentResponse, error)
	Update(ctx context.Context, req *ct.UpdateEventStudentRequest) (*ct.GetEventStudentResponse, error)
	Delete(ctx context.Context, req *ct.EventStudentID) (*ct.EventStudentEmpty, error)
    GetStudentWithEventsByID(ctx context.Context, req *ct.StudentID) (*ct.GetStudentWithEventsResponse, error)
}

type StudentPaymentRepoI interface {
    Create(ctx context.Context, req *ct.CreateStudentPaymentRequest) (*ct.StudentPaymentResponse, error)
    GetByID(ctx context.Context, req *ct.StudentPaymentID) (*ct.GetStudentPaymentResponse, error)
    GetList(ctx context.Context, req *ct.GetListStudentPaymentRequest) (*ct.GetListStudentPaymentResponse, error)
    Update(ctx context.Context, req *ct.UpdateStudentPaymentRequest) (*ct.GetStudentPaymentResponse, error)
    Delete(ctx context.Context, req *ct.StudentPaymentID) (*ct.StudentPaymentEmpty, error)
}