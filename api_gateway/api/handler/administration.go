package handler

import (
	"api_gateway/genproto/user_service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/administration [post]
// @Summary Create administration
// @Description API for creating an administration
// @Tags administration
// @Accept json
// @Produce json
// @Param administration body user_service.CreateAdministrationRequest true "Administration"
// @Success 200 {object} user_service.AdministrationResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateAdministration(c *gin.Context) {
	var req user_service.CreateAdministrationRequest
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not a Manager")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.AdministrationService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /administration/{id} [get]
// @Summary Get a single administration by ID
// @Description API for getting a single administration by ID
// @Tags administration
// @Accept json
// @Produce json
// @Param id path string true "administration ID"
// @Success 200 {object} user_service.GetAdministrationResponse
// @Failure 404 {object} models.ResponseError "Administration not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetAdministrationByID(c *gin.Context) {
	id := c.Param("id")
	
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting auth")
		return
	}

	if data.UserRole != "SuperAdmin" && data.UserRole != "Manager" && data.UserRole != "Administration" {
		handleGrpcErrWithDescription(c, h.log, errors.New("Unauthorized"), "You are not authorized")
		return
	}

	if data.UserRole == "Administration" {
		id = data.UserID
	}
	req := &user_service.AdministrationID{Id: id}

	resp, err := h.grpcClient.AdministrationService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /administration/{id} [put]
// @Summary Update an administration by ID
// @Description API for updating an administration by ID
// @Tags administration
// @Accept json
// @Produce json
// @Param id path string true "administration ID"
// @Param administration body user_service.UpdateAdministrationRequest true "Administration"
// @Success 200 {object} user_service.GetAdministrationResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Administration not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateAdministration(c *gin.Context) {
	id := c.Param("id")
	var req user_service.UpdateAdministrationRequest
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
	resp, err := h.grpcClient.AdministrationService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /administration/list [get]
// @Summary Get List of Administrations
// @Description API for getting a list of administrations
// @Tags administration
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListAdministrationResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListAdministration(c *gin.Context) {
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

	req := user_service.GetListAdministrationRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.AdministrationService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /administration/{id} [delete]
// @Summary Delete an administration by ID
// @Description API for deleting an administration by ID
// @Tags administration
// @Accept json
// @Produce json
// @Param id path string true "administration ID"
// @Success 200 {object} user_service.AdministrationEmpty
// @Failure 404 {object} models.ResponseError "Administration not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteAdministration(c *gin.Context) {
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
	req := &user_service.AdministrationID{Id: id}

	resp, err := h.grpcClient.AdministrationService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}



// @Security ApiKeyAuth
// @Router /administration/report/list [get]
// @Summary Get List of Administrations
// @Description API for getting a list of administrations
// @Tags Report
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetReportListAdministrationResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetReportListAdministration(c *gin.Context) {
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

	req := user_service.GetReportListAdministrationRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.AdministrationService().GetReportList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
