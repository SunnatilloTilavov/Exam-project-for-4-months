package api

import (
	"net/http"

	_ "api_gateway/api/docs" //for swagger
	"api_gateway/api/handler"
	"api_gateway/config"
	"api_gateway/pkg/grpc_client"
	"api_gateway/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Bu xavfsiz emas, kerakli domenlarni qo'shishingiz kerak
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	// Teacher endpoints
	r.POST("/create/teacher", handler.CreateTeacher)
	r.GET("/teacher/:id", handler.GetTeacherByID)
	r.GET("/teacher/list", handler.GetListTeacher)
	r.PUT("/teacher/:id", handler.UpdateTeacher)
	r.DELETE("/teacher/:id", handler.DeleteTeacher)

	// SupportTeacher endpoints
	r.POST("/create/supportteacher", handler.CreateSupportTeacher)
	r.GET("/supportteacher/:id", handler.GetSupportTeacherByID)
	r.GET("/supportteacher/list", handler.GetListSupportTeacher)
	r.PUT("/supportteacher/:id", handler.UpdateSupportTeacher)
	r.DELETE("/supportteacher/:id", handler.DeleteSupportTeacher)

	// Student endpoints
	r.POST("/create/student", handler.CreateStudent)
	r.GET("/student/:id", handler.GetStudentByID)
	r.GET("/student/list", handler.GetListStudent)
	r.PUT("/student/:id", handler.UpdateStudent)
	r.DELETE("/student/:id", handler.DeleteStudent)

	// Manager endpoints
	r.POST("/create/manager", handler.CreateManager)
	r.GET("/manager/:id", handler.GetManagerByID)
	r.GET("/manager/list", handler.GetListManager)
	r.PUT("/manager/:id", handler.UpdateManager)
	r.DELETE("/manager/:id", handler.DeleteManager)

	// Administration endpoints
	r.POST("/create/administration", handler.CreateAdministration)
	r.GET("/administration/:id", handler.GetAdministrationByID)
	r.GET("/administration/list", handler.GetListAdministration)
	r.PUT("/administration/:id", handler.UpdateAdministration)
	r.DELETE("/administration/:id", handler.DeleteAdministration)

	// Login endpoints
	r.POST("/login/administration", handler.AdministarationLogin)
	r.POST("/login/manager", handler.ManagerLogin)
	r.POST("/login/student", handler.StudentLogin)
	r.POST("/login/support-teacher", handler.SupportTeacherLogin)
	r.POST("/login/teacher", handler.TeacherLogin)
	r.POST("/login/superadmin", handler.SuperAdminLogin)

	// Branch endpoints
	r.POST("/create/branch", handler.CreateBranch)
	r.GET("/branch/:id", handler.GetBranchByID)
	r.GET("/branch/list", handler.GetListBranch)
	r.PUT("/branch/:id", handler.UpdateBranch)
	r.DELETE("/branch/:id", handler.DeleteBranch)

	// Event Student endpoints
	r.POST("/create/event_student", handler.CreateEventStudent)
	r.GET("/event_student/:id", handler.GetEventStudentByID)
	r.GET("/event_student/list", handler.GetListEventStudent)
	r.PUT("/event_student/:id", handler.UpdateEventStudent)
	r.DELETE("/event_student/:id", handler.DeleteEventStudent)
	r.GET("/event_student/student/:id", handler.GetStudentWithEventsByID)

	// Event endpoints
	r.POST("/create/event", handler.CreateEvent)
	r.GET("/event/:id", handler.GetEventByID)
	r.GET("/event/list", handler.GetListEvent)
	r.PUT("/event/:id", handler.UpdateEvent)
	r.DELETE("/event/:id", handler.DeleteEvent)

	// Jurnals endpoints
	r.POST("/api/v1/jurnals", handler.CreateJurnal)
	r.GET("/api/v1/jurnals/:id", handler.GetJurnalByID)
	r.GET("/api/v1/jurnals/student/:id", handler.GetJurnalByIDStudent)
	r.GET("/api/v1/jurnals", handler.GetListJurnal)
	r.PUT("/api/v1/jurnals/:id", handler.UpdateJurnal)
	r.DELETE("/api/v1/jurnals/:id", handler.DeleteJurnal)

	// Schedule endpoints
	r.POST("/api/v1/schedules", handler.CreateSchedule)
	r.GET("/api/v1/schedules/:id", handler.GetScheduleByID)
	r.GET("/api/v1/schedules/list", handler.GetListSchedule)
	r.PUT("/api/v1/schedules/:id", handler.UpdateSchedule)
	r.DELETE("/api/v1/schedules/:id", handler.DeleteSchedule)

	// Group endpoints
	r.POST("/api/v1/groups", handler.CreateGroup)
	r.GET("/api/v1/groups/:id", handler.GetGroupByID)
	r.GET("/api/v1/groups/teacher/:id", handler.GetGroupByIDTeacher)
	r.GET("/api/v1/groups", handler.GetListGroup)
	r.PUT("/api/v1/groups/:id", handler.UpdateGroup)
	r.DELETE("/api/v1/groups/:id", handler.DeleteGroup)

	//Student Payment endpoints
	r.POST("/create/student_payment", handler.CreateStudentPayment)
	r.GET("/student_payment/:id", handler.GetStudentPaymentByID)
	r.GET("/student_payment/list", handler.GetListStudentPayment)
	r.PUT("/student_payment/:id", handler.UpdateStudentPayment)
	r.DELETE("/student_payment/:id", handler.DeleteStudentPayment)

	//Student task endpoints
	r.POST("/create/student_task", handler.CreateStudentTask)
	r.GET("/student_task/:id", handler.GetStudentTaskByID)
	r.GET("/student_task/student/:id", handler.GetStudentTaskByStudentID)
	r.GET("/student_task/list", handler.GetListStudentTasks)
	r.PUT("/student_task/:id", handler.UpdateStudentTask)
	r.DELETE("/student_task/:id", handler.DeleteStudentTask)
	r.PUT("/student_task/score/teacher/:id", handler.UpdateScoreForTeacher)
	r.PUT("/student_task/score/student/:id", handler.UpdateScoreForStudent)

	// Task endpoints
	r.POST("/create/task", handler.CreateTask)
	r.GET("/task/:id", handler.GetTaskByID)
	r.GET("/task/list", handler.GetListTask)
	r.PUT("/task/:id", handler.UpdateTask)
	r.DELETE("/task/:id", handler.DeleteTask)

	//Report endpoints
	r.GET("/administration/report/list", handler.GetReportListAdministration)
	r.GET("/teacher/report/list", handler.GetReportListTeacher)
	r.GET("/supportteacher/report/list", handler.GetReportListSupportTeacher)
	r.GET("/student/report/list", handler.GetReportListStudent)
	
	// Swagger endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
