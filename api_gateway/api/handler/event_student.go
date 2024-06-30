package handler

import (
	"api_gateway/genproto/education_management_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/event_student [post]
// @Summary Create event student
// @Description API for creating an event student
// @Tags event_student
// @Accept json
// @Produce json
// @Param event_student body education_management_service.CreateEventStudentRequest true "Event Student"
// @Success 200 {object} education_management_service.EventStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateEventStudent(c *gin.Context) {
	var req education_management_service.CreateEventStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.EventStudentService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event_student/{id} [get]
// @Summary Get event student by ID
// @Description API for getting an event student by ID
// @Tags event_student
// @Accept json
// @Produce json
// @Param id path string true "Event Student ID"
// @Success 200 {object} education_management_service.GetEventStudentResponse
// @Failure 404 {object} models.ResponseError "Event Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetEventStudentByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.EventStudentID{Id: id}

	resp, err := h.grpcClient.EventStudentService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event_student/list [get]
// @Summary Get list of event students
// @Description API for getting a list of event students
// @Tags event_student
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListEventStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListEventStudent(c *gin.Context) {
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

	req := education_management_service.GetListEventStudentRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.EventStudentService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event_student/{id} [put]
// @Summary Update event student by ID
// @Description API for updating an event student by ID
// @Tags event_student
// @Accept json
// @Produce json
// @Param id path string true "Event Student ID"
// @Param event_student body education_management_service.UpdateEventStudentRequest true "Event Student"
// @Success 200 {object} education_management_service.GetEventStudentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Event Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateEventStudent(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateEventStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.EventStudentService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event_student/{id} [delete]
// @Summary Delete event student by ID
// @Description API for deleting an event student by ID
// @Tags event_student
// @Accept json
// @Produce json
// @Param id path string true "Event Student ID"
// @Success 200 {object} education_management_service.EventStudentEmpty
// @Failure 404 {object} models.ResponseError "Event Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteEventStudent(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.EventStudentID{Id: id}

	resp, err := h.grpcClient.EventStudentService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /event_student/student/{id} [get]
// @Summary Get student with events by student ID
// @Description API for getting a student with their events by student ID
// @Tags event_student
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} education_management_service.GetStudentWithEventsResponse
// @Failure 404 {object} models.ResponseError "Student not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetStudentWithEventsByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.StudentID{Id: id}

	resp, err := h.grpcClient.EventStudentService().GetStudentByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
