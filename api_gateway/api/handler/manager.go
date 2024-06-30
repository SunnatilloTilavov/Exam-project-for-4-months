package handler

import (
	"api_gateway/genproto/user_service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/manager [post]
// @Summary Create manager
// @Description API for creating a manager
// @Tags manager
// @Accept json
// @Produce json
// @Param manager body user_service.CreateManagerRequest true "Manager"
// @Success 200 {object} user_service.ManagerResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateManager(c *gin.Context) {
	var req user_service.CreateManagerRequest
	data, err := getAuthInfo(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while getting auth")
		return
	}

	if data.UserRole!="SuperAdmin"||data.UserID!="cb618114-b1ba-4382-8cdb-8002e3a4aa48"{
		handleGrpcErrWithDescription(c, h.log,err, "You are not a SUPER admin")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.ManagerService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}




// @Security ApiKeyAuth
// @Router /manager/{id} [get]
// @Summary Get a single manager by ID
// @Description API for getting a single manager by ID
// @Tags manager
// @Accept json
// @Produce json
// @Param id path string true "manager ID"
// @Success 200 {object} user_service.GetManagerResponse
// @Failure 404 {object} models.ResponseError "Manager not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetManagerByID(c *gin.Context) {
	id := c.Param("id")
	req := &user_service.ManagerID{Id: id}

	resp, err := h.grpcClient.ManagerService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /manager/{id} [put]
// @Summary Update a manager by ID
// @Description API for updating a manager by ID
// @Tags manager
// @Accept json
// @Produce json
// @Param id path string true "manager ID"
// @Param manager body user_service.UpdateManagerRequest true "Manager"
// @Success 200 {object} user_service.GetManagerResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Manager not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateManager(c *gin.Context) {
	id := c.Param("id")
	var req user_service.UpdateManagerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.ManagerService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /manager/list [get]
// @Summary Get List of Managers
// @Description API for getting a list of managers
// @Tags manager
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} user_service.GetListManagerResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListManager(c *gin.Context) {
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

	req := user_service.GetListManagerRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.ManagerService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /manager/{id} [delete]
// @Summary Delete a manager by ID
// @Description API for deleting a manager by ID
// @Tags manager
// @Accept json
// @Produce json
// @Param id path string true "manager ID"
// @Success 200 {object} user_service.ManagerEmpty
// @Failure 404 {object} models.ResponseError "Manager not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteManager(c *gin.Context) {
	id := c.Param("id")
	req := &user_service.ManagerID{Id: id}

	resp, err := h.grpcClient.ManagerService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
