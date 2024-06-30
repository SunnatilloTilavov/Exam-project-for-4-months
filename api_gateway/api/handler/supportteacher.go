package handler

import (
	"api_gateway/genproto/user_service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/supportteacher [post]
// @Summary Create support teacher
// @Description API for creating a support teacher
// @Tags support_teacher
// @Accept json
// @Produce json
// @Param support_teacher body user_service.CreateSupportTeacherRequest true "Support Teacher"
// @Success 200 {object} user_service.SupportTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateSupportTeacher(c *gin.Context) {
	var req user_service.CreateSupportTeacherRequest
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

	resp, err := h.grpcClient.SupportTeacherService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /supportteacher/{id} [get]
// @Summary Get a single support teacher by ID
// @Description API for getting a single support teacher by ID
// @Tags support_teacher
// @Accept json
// @Produce json
// @Param id path string true "support teacher ID"
// @Success 200 {object} user_service.GetSupportTeacherResponse
// @Failure 404 {object} models.ResponseError "Support Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetSupportTeacherByID(c *gin.Context) {
	id := c.Param("id")

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" && data.UserRole != "SupportTeacher" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	if data.UserRole == "SupportTeacher" {
		id = data.UserID
	}

	req := &user_service.SupportTeacherID{Id: id}

	resp, err := h.grpcClient.SupportTeacherService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /supportteacher/{id} [put]
// @Summary Update a support teacher by ID
// @Description API for updating a support teacher by ID
// @Tags support_teacher
// @Accept json
// @Produce json
// @Param id path string true "support teacher ID"
// @Param support_teacher body user_service.UpdateSupportTeacherRequest true "Support Teacher"
// @Success 200 {object} user_service.GetSupportTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Support Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateSupportTeacher(c *gin.Context) {
	id := c.Param("id")

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	var req user_service.UpdateSupportTeacherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.SupportTeacherService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /supportteacher/list [get]
// @Summary Get List of Support Teachers
// @Description API for getting a list of support teachers
// @Tags support_teacher
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListSupportTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListSupportTeacher(c *gin.Context) {
	limitStr := c.Query("limit")

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

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

	req := user_service.GetListSupportTeacherRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.SupportTeacherService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /supportteacher/{id} [delete]
// @Summary Delete a support teacher by ID
// @Description API for deleting a support teacher by ID
// @Tags support_teacher
// @Accept json
// @Produce json
// @Param id path string true "support teacher ID"
// @Success 200 {object} user_service.SupportTeacherEmpty
// @Failure 404 {object} models.ResponseError "Support Teacher not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteSupportTeacher(c *gin.Context) {
	id := c.Param("id")
	req := &user_service.SupportTeacherID{Id: id}

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}




// @Security ApiKeyAuth
// @Router /supportteacher/report/list [get]
// @Summary Get List of Support Teachers
// @Description API for getting a list of support teachers
// @Tags Report
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListSupportTeacherResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetReportListSupportTeacher(c *gin.Context) {
	limitStr := c.Query("limit")

	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

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

	req := user_service.GetReportListSupportTeacherRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.SupportTeacherService().GetReportList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

