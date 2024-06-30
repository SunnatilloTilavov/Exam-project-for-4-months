package handler

import (
	"api_gateway/genproto/user_service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/student [post]
// @Summary Create student
// @Description API for creating a student
// @Tags student
// @Accept json
// @Produce json
// @Param student body user_service.CreateStudentRequest true "Student"
// @Success 200 {object} user_service.StudentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateStudent(c *gin.Context) {
	var req user_service.CreateStudentRequest
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.StudentService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student/{id} [get]
// @Summary Get a single student by ID
// @Description API for getting a single student by ID
// @Tags student
// @Accept json
// @Produce json
// @Param id path string true "student ID"
// @Success 200 {object} user_service.GetStudentResponse
// @Failure 404 {object} models.ResponseError "Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" && data.UserRole != "Student" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	if data.UserRole == "Student" {
		id = data.UserID
	}

	req := &user_service.StudentID{Id: id}

	resp, err := h.grpcClient.StudentService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student/{id} [put]
// @Summary Update a student by ID
// @Description API for updating a student by ID
// @Tags student
// @Accept json
// @Produce json
// @Param id path string true "student ID"
// @Param student body user_service.UpdateStudentRequest true "Student"
// @Success 200 {object} user_service.GetStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var req user_service.UpdateStudentRequest
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.StudentService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student/list [get]
// @Summary Get List of Students
// @Description API for getting a list of students
// @Tags student
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListStudent(c *gin.Context) {
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}
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

	req := user_service.GetListStudentRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.StudentService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student/{id} [delete]
// @Summary Delete a student by ID
// @Description API for deleting a student by ID
// @Tags student
// @Accept json
// @Produce json
// @Param id path string true "student ID"
// @Success 200 {object} user_service.StudentEmpty
// @Failure 404 {object} models.ResponseError "Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteStudent(c *gin.Context) {
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}
	id := c.Param("id")
	req := &user_service.StudentID{Id: id}

	resp, err := h.grpcClient.StudentService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Security ApiKeyAuth
// @Router /student/report/list [get]
// @Summary Get List of Students
// @Description API for getting a list of students
// @Tags Report
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetReportListStudent(c *gin.Context) {
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin"  {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}
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

	req := user_service.GetReportListStudentRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.StudentService().GetReportList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}