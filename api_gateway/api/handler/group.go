package handler

import (
	"api_gateway/genproto/education_management_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateGroup handles the RPC CreateGroup
// @Summary Create a new Group
// @Description Create a new Group entry
// @Accept json
// @Tags Group
// @Produce json
// @Param request body education_management_service.CreateGroupRequest true "Group creation request"
// @Success 200 {object} education_management_service.GroupResponse
// @Router /api/v1/groups [post]
func (h *handler) CreateGroup(c *gin.Context) {
	var req education_management_service.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.GroupService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetGroupByID handles the RPC GetByID
// @Summary Get a Group by ID
// @Description Get a Group entry by its ID
// @Produce json
// @Tags Group
// @Param id path string true "Group ID"
// @Success 200 {object} education_management_service.GetGroupResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/groups/{id} [get]
func (h *handler) GetGroupByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.GroupID{Id: id}

	resp, err := h.grpcClient.GroupService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetGroupByIDTeacher handles the RPC GetByIDTeacher
// @Summary Get Groups by Teacher ID
// @Description Get Groups associated with a Teacher by Teacher ID
// @Produce json
// @Tags Group
// @Param id path string true "Teacher ID"
// @Success 200 {object} education_management_service.GetListGroupResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/groups/teacher/{id} [get]
func (h *handler) GetGroupByIDTeacher(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.TeacherID{Id: id}

	resp, err := h.grpcClient.GroupService().GetByIDTeacher(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListGroup handles the RPC GetList
// @Summary Get a list of Groups
// @Description Get a list of Group entries with pagination and search
// @Produce json
// @Tags Group
// @Param limit query int false "Limit per page"
// @Param page query int false "Page number"
// @Param search query string false "Search query"
// @Success 200 {object} education_management_service.GetListGroupResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/groups [get]
func (h *handler) GetListGroup(c *gin.Context) {
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

	req := education_management_service.GetListGroupRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.GroupService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateGroup handles the RPC UpdateGroup
// @Summary Update a Group by ID
// @Description Update a Group entry by its ID
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param request body education_management_service.UpdateGroupRequest true "Group update request"
// @Success 200 {object} education_management_service.GetGroupResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/groups/{id} [put]
func (h *handler) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.GroupService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteGroup handles the RPC Delete
// @Summary Delete a Group by ID
// @Description Delete a Group entry by its ID
// @Produce json
// @Tags Group
// @Param id path string true "Group ID"
// @Success 200 {object} education_management_service.GroupEmpty
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/groups/{id} [delete]
func (h *handler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.GroupID{Id: id}

	resp, err := h.grpcClient.GroupService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
