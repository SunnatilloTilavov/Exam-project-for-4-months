package handler

import (
	"api_gateway/genproto/education_management_service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/branch [post]
// @Summary Create branch
// @Description API for creating a branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch body education_management_service.CreateBranchRequest true "Branch"
// @Success 200 {object} education_management_service.BranchResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateBranch(c *gin.Context) {
	var req education_management_service.CreateBranchRequest
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

	resp, err := h.grpcClient.BranchService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /branch/{id} [get]
// @Summary Get a single branch by ID
// @Description API for getting a single branch by ID
// @Tags branch
// @Accept json
// @Produce json
// @Param id path string true "branch ID"
// @Success 200 {object} education_management_service.GetBranchResponse
// @Failure 404 {object} models.ResponseError "Branch not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetBranchByID(c *gin.Context) {
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

	req := &education_management_service.BranchID{Id: id}

	resp, err := h.grpcClient.BranchService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /branch/{id} [put]
// @Summary Update a branch by ID
// @Description API for updating a branch by ID
// @Tags branch
// @Accept json
// @Produce json
// @Param id path string true "branch ID"
// @Param branch body education_management_service.UpdateBranchRequest true "Branch"
// @Success 200 {object} education_management_service.GetBranchResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Branch not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateBranch(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateBranchRequest
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
	resp, err := h.grpcClient.BranchService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /branch/list [get]
// @Summary Get List of Branches
// @Description API for getting a list of branches
// @Tags branch
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListBranchResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListBranch(c *gin.Context) {
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

	req := education_management_service.GetListBranchRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.BranchService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /branch/{id} [delete]
// @Summary Delete a branch by ID
// @Description API for deleting a branch by ID
// @Tags branch
// @Accept json
// @Produce json
// @Param id path string true "branch ID"
// @Success 200 {object} education_management_service.BranchEmpty
// @Failure 404 {object} models.ResponseError "Branch not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteBranch(c *gin.Context) {
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
	req := &education_management_service.BranchID{Id: id}

	resp, err := h.grpcClient.BranchService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
