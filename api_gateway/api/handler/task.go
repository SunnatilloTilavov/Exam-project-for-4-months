package handler

import (
	"api_gateway/genproto/education_management_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/task [post]
// @Summary Create task
// @Description API for creating a task
// @Tags task
// @Accept json
// @Produce json
// @Param task body education_management_service.CreateTaskRequest true "Task"
// @Success 200 {object} education_management_service.TaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateTask(c *gin.Context) {
	var req education_management_service.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.TaskService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /task/{id} [get]
// @Summary Get task by ID
// @Description API for getting a task by ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} education_management_service.GetTaskResponse
// @Failure 404 {object} models.ResponseError "Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.TaskID{Id: id}

	resp, err := h.grpcClient.TaskService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /task/list [get]
// @Summary Get list of tasks
// @Description API for getting a list of tasks
// @Tags task
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListTask(c *gin.Context) {
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

	req := education_management_service.GetListTaskRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.TaskService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /task/{id} [put]
// @Summary Update task by ID
// @Description API for updating a task by ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body education_management_service.UpdateTaskRequest true "Task"
// @Success 200 {object} education_management_service.GetTaskResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.TaskService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /task/{id} [delete]
// @Summary Delete task by ID
// @Description API for deleting a task by ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} education_management_service.TaskEmpty
// @Failure 404 {object} models.ResponseError "Task not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.TaskID{Id: id}

	resp, err := h.grpcClient.TaskService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
