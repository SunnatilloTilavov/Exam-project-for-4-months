package handler

import (
	"api_gateway/genproto/user_service"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/teacher [post]
// @Summary Create teacher
// @Description API for creating a teacher
// @Tags teacher
// @Accept json
// @Produce json
// @Param teacher body user_service.CreateTeacherRequest true "Teacher"
// @Success 200 {object} user_service.TeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateTeacher(c *gin.Context) {
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}
	var req user_service.CreateTeacherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.TeacherService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /teacher/{id} [get]
// @Summary Get a single teacher by ID
// @Description API for getting a single teacher by ID
// @Tags teacher
// @Accept json
// @Produce json
// @Param id path string true "teacher ID"
// @Success 200 {object} user_service.GetTeacherResponse
// @Failure 404 {object} models.ResponseError "Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" && data.UserRole != "Teacher" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}
	
	if data.UserRole == "Teacher" {
		id = data.UserID
	}

	req := &user_service.TeacherID{Id: id}

	resp, err := h.grpcClient.TeacherService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /teacher/{id} [put]
// @Summary Update a teacher by ID
// @Description API for updating a teacher by ID
// @Tags teacher
// @Accept json
// @Produce json
// @Param id path string true "teacher ID"
// @Param teacher body user_service.UpdateTeacherRequest true "Teacher"
// @Success 200 {object} user_service.GetTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateTeacher(c *gin.Context) {
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
	var req user_service.UpdateTeacherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.TeacherService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /teacher/list [get]
// @Summary Get List of Teachers
// @Description API for getting a list of teachers
// @Tags teacher
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListTeacher(c *gin.Context) {
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

	req := user_service.GetListTeacherRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.TeacherService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /teacher/{id} [delete]
// @Summary Delete a teacher by ID
// @Description API for deleting a teacher by ID
// @Tags teacher
// @Accept json
// @Produce json
// @Param id path string true "teacher ID"
// @Success 200 {object} user_service.TeacherEmpty
// @Failure 404 {object} models.ResponseError "Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteTeacher(c *gin.Context) {
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
	req := &user_service.TeacherID{Id: id}

	resp, err := h.grpcClient.TeacherService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Security ApiKeyAuth
// @Router /teacher/report/list [get]
// @Summary Get List of Teachers
// @Description API for getting a list of teachers
// @Tags Report
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetReportListTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetReportListTeacher(c *gin.Context) {
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" {
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

	req := user_service.GetReportListTeacherRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.TeacherService().GetReportList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
