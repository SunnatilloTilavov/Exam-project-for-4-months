package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api_gateway/genproto/education_management_service"
)
// @Summary Get student task by ID
// @Description API for getting a student task by ID
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student Task ID"
// @Success 200 {object} education_management_service.GetStudentTaskResponse
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/{id} [get]
func (h *handler) GetStudentTaskByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.StudentTaskID{Id: id}

	resp, err := h.grpcClient.StudentTaskService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get list of student tasks
// @Description API for getting a list of student tasks
// @Tags student_task
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListStudentTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/list [get]
func (h *handler) GetListStudentTasks(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid limit parameter")
		return
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid page parameter")
		return
	}

	req := education_management_service.GetListStudentTaskRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.StudentTaskService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Create student task
// @Description API for creating a student task
// @Tags student_task
// @Accept json
// @Produce json
// @Param student_task body education_management_service.CreateStudentTaskRequest true "Student Task"
// @Success 200 {object} education_management_service.StudentTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /create/student_task [post]
func (h *handler) CreateStudentTask(c *gin.Context) {
	var req education_management_service.CreateStudentTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.StudentTaskService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update student task by ID
// @Description API for updating a student task by ID
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student Task ID"
// @Param student_task body education_management_service.UpdateStudentTaskRequest true "Student Task"
// @Success 200 {object} education_management_service.GetStudentTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/{id} [put]
func (h *handler) UpdateStudentTask(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateStudentTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.StudentTaskService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Delete student task by ID
// @Description API for deleting a student task by ID
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student Task ID"
// @Success 200 {object} education_management_service.StudentTaskEmpty
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/{id} [delete]
func (h *handler) DeleteStudentTask(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.StudentTaskID{Id: id}

	resp, err := h.grpcClient.StudentTaskService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get student task by student ID
// @Description API for getting a student task by student ID
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} education_management_service.GetStudentTaskResponse
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/student/{id} [get]
func (h *handler) GetStudentTaskByStudentID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.TaskStudentID{Id: id}

	resp, err := h.grpcClient.StudentTaskService().GetByIDStudent(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update student task score for teacher
// @Description API for updating a student task score for a teacher
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student Task ID"
// @Param score body education_management_service.UpdateStudentScoreRequest true "Student Task Score"
// @Success 200 {object} education_management_service.GetStudentTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/score/teacher/{id} [put]
func (h *handler) UpdateScoreForTeacher(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateStudentScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.StudentTaskService().UpdateScoreforTeacher(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update student task score for student
// @Description API for updating a student task score for a student
// @Tags student_task
// @Accept json
// @Produce json
// @Param id path string true "Student Task ID"
// @Param score body education_management_service.UpdateStudentScoreRequest true "Student Task Score"
// @Success 200 {object} education_management_service.GetStudentTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Student Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /student_task/score/student/{id} [put]
func (h *handler) UpdateScoreForStudent(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateStudentScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.StudentTaskService().UpdateScoreforStudent(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}