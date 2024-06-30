package storage

import (
	"context"
	ct "user_service/genproto/user_service"
	// "github.com/google/uuid"
)

type StorageI interface {
	CloseDB()
	Teacher() TeacherRepoI
	SupportTeacher() SupportTeacherRepoI
	Administration() AdministrationRepoI
	Student() StudentRepoI
	Manager() ManagerRepoI
}

type TeacherRepoI interface {
	Create(ctx context.Context, req *ct.CreateTeacherRequest) (*ct.TeacherResponse, error)
	GetByID(ctx context.Context, req *ct.TeacherID) (*ct.GetTeacherResponse, error)
	GetList(ctx context.Context, req *ct.GetListTeacherRequest) (*ct.GetListTeacherResponse, error)
	Update(ctx context.Context, req *ct.UpdateTeacherRequest) (*ct.GetTeacherResponse, error)
	GetByLogin(ctx context.Context, login string) (*ct.GetTeacherResponse, error)
	Delete(ctx context.Context, req *ct.TeacherID) (*ct.TeacherEmpty, error)
	GetReportList(ctx context.Context, req *ct.GetReportListTeacherRequest) (*ct.GetReportListTeacherResponse, error)
}

type SupportTeacherRepoI interface {
	Create(ctx context.Context, req *ct.CreateSupportTeacherRequest) (*ct.SupportTeacherResponse, error)
	GetByID(ctx context.Context, req *ct.SupportTeacherID) (*ct.GetSupportTeacherResponse, error)
	GetList(ctx context.Context, req *ct.GetListSupportTeacherRequest) (*ct.GetListSupportTeacherResponse, error)
	Update(ctx context.Context, req *ct.UpdateSupportTeacherRequest) (*ct.GetSupportTeacherResponse, error)
	Delete(ctx context.Context, req *ct.SupportTeacherID) (*ct.SupportTeacherEmpty, error)
	GetByLogin(ctx context.Context, login string) (*ct.GetSupportTeacherResponse, error)
	GetReportList(ctx context.Context, req *ct.GetReportListSupportTeacherRequest) (*ct.GetReportListSupportTeacherResponse, error) 
}

type AdministrationRepoI interface {
	Create(ctx context.Context, req *ct.CreateAdministrationRequest) (*ct.AdministrationResponse, error)
	GetByID(ctx context.Context, req *ct.AdministrationID) (*ct.GetAdministrationResponse, error)
	GetList(ctx context.Context, req *ct.GetListAdministrationRequest) (*ct.GetListAdministrationResponse, error)
	Update(ctx context.Context, req *ct.UpdateAdministrationRequest) (*ct.GetAdministrationResponse, error)
	GetByLogin(ctx context.Context, login string) (*ct.GetAdministrationResponse, error)
	Delete(ctx context.Context, req *ct.AdministrationID) (*ct.AdministrationEmpty, error)
	GetReportList(ctx context.Context, req *ct.GetReportListAdministrationRequest) (*ct.GetReportListAdministrationResponse, error)
}

type StudentRepoI interface {
	Create(ctx context.Context, req *ct.CreateStudentRequest) (*ct.StudentResponse, error)
	GetByID(ctx context.Context, req *ct.StudentID) (*ct.GetStudentResponse, error)
	GetList(ctx context.Context, req *ct.GetListStudentRequest) (*ct.GetListStudentResponse, error)
	Update(ctx context.Context, req *ct.UpdateStudentRequest) (*ct.GetStudentResponse, error)
	GetByLogin(ctx context.Context, login string) (*ct.GetStudentResponse, error)
	Delete(ctx context.Context, req *ct.StudentID) (*ct.StudentEmpty, error)
	GetReportList(ctx context.Context, req *ct.GetReportListStudentRequest) (*ct.GetReportListStudentResponse, error)
}

type ManagerRepoI interface {
	Create(ctx context.Context, req *ct.CreateManagerRequest) (*ct.ManagerResponse, error)
	GetByID(ctx context.Context, req *ct.ManagerID) (*ct.GetManagerResponse, error)
	GetList(ctx context.Context, req *ct.GetListManagerRequest) (*ct.GetListManagerResponse, error)
	Delete(ctx context.Context, req *ct.ManagerID) (*ct.ManagerEmpty, error)
	Update(ctx context.Context, req *ct.UpdateManagerRequest) (*ct.GetManagerResponse, error)
	GetByLogin(ctx context.Context, login string) (*ct.GetManagerResponse, error)
}
