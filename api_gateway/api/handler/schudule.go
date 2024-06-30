package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api_gateway/genproto/education_management_service"
)

// CreateSchedule handles the RPC CreateSchedule
// @Summary Create a new Schedule
// @Description Create a new Schedule entry
// @Accept json
// @Tags Schedule
// @Produce json
// @Param request body education_management_service.CreateScheduleRequest true "Schedule creation request"
// @Success 200 {object} education_management_service.ScheduleResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/schedules [post]
func (h *handler) CreateSchedule(c *gin.Context) {
	var req education_management_service.CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.ScheduleService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetScheduleByID handles the RPC GetByID
// @Summary Get a Schedule by ID
// @Description Get a Schedule entry by its ID
// @Produce json
// @Tags Schedule
// @Param id path string true "Schedule ID"
// @Success 200 {object} education_management_service.GetScheduleResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/schedules/{id} [get]
func (h *handler) GetScheduleByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.ScheduleID{Id: id}

	resp, err := h.grpcClient.ScheduleService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListSchedule handles the RPC GetList
// @Summary Get a list of Schedules
// @Description Get a list of Schedule entries with pagination and search
// @Produce json
// @Tags Schedule
// @Param limit query int false "Limit per page"
// @Param page query int false "Page number"
// @Param search query string false "Search query"
// @Success 200 {object} education_management_service.GetListScheduleResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/schedules [get]
func (h *handler) GetListSchedule(c *gin.Context) {
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

	req := education_management_service.GetListScheduleRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.ScheduleService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSchedule handles the RPC UpdateSchedule
// @Summary Update a Schedule by ID
// @Description Update a Schedule entry by its ID
// @Accept json
// @Produce json
// @Tags Schedule
// @Param id path string true "Schedule ID"
// @Param request body education_management_service.UpdateScheduleRequest true "Schedule update request"
// @Success 200 {object} education_management_service.GetScheduleResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/schedules/{id} [put]
func (h *handler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.ScheduleService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteSchedule handles the RPC Delete
// @Summary Delete a Schedule by ID
// @Description Delete a Schedule entry by its ID
// @Produce json
// @Tags Schedule
// @Param id path string true "Schedule ID"
// @Success 200 {object} education_management_service.ScheduleEmpty
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/schedules/{id} [delete]
func (h *handler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.ScheduleID{Id: id}

	resp, err := h.grpcClient.ScheduleService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
