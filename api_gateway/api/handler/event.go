package handler

import (
	"api_gateway/genproto/education_management_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/event [post]
// @Summary Create event
// @Description API for creating an event
// @Tags event
// @Accept json
// @Produce json
// @Param event body education_management_service.CreateEventRequest true "Event"
// @Success 200 {object} education_management_service.EventResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateEvent(c *gin.Context) {
	var req education_management_service.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.EventService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event/{id} [get]
// @Summary Get event by ID
// @Description API for getting an event by ID
// @Tags event
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} education_management_service.GetEventResponse
// @Failure 404 {object} models.ResponseError "Event not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetEventByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.EventID{Id: id}

	resp, err := h.grpcClient.EventService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event/list [get]
// @Summary Get list of events
// @Description API for getting a list of events
// @Tags event
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListEventResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListEvent(c *gin.Context) {
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

	req := education_management_service.GetListEventRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.EventService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event/{id} [put]
// @Summary Update event by ID
// @Description API for updating an event by ID
// @Tags event
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param event body education_management_service.UpdateEventRequest true "Event"
// @Success 200 {object} education_management_service.GetEventResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Event not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.EventService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event/{id} [delete]
// @Summary Delete event by ID
// @Description API for deleting an event by ID
// @Tags event
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} education_management_service.EventEmpty
// @Failure 404 {object} models.ResponseError "Event not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.EventID{Id: id}

	resp, err := h.grpcClient.EventService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
