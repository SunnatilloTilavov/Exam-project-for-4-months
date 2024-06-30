package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api_gateway/genproto/education_management_service"
)

// CreateJurnal handles the RPC CreateJurnal
// @Summary Create a new Jurnal
// @Description Create a new Jurnal entry
// @Accept json
// @Tags Jurnal
// @Produce json
// @Param request body education_management_service.CreateJurnalRequest true "Jurnal creation request"
// @Success 200 {object} education_management_service.JurnalResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals [post]
func (h *handler) CreateJurnal(c *gin.Context) {
	var req education_management_service.CreateJurnalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.JurnalService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetJurnalByID handles the RPC GetByID
// @Summary Get a Jurnal by ID
// @Description Get a Jurnal entry by its ID
// @Produce json
// @Tags Jurnal
// @Param id path string true "Jurnal ID"
// @Success 200 {object} education_management_service.GetJurnalResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals/{id} [get]
func (h *handler) GetJurnalByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.JurnalID{Id: id}

	resp, err := h.grpcClient.JurnalService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetJurnalByIDStudent handles the RPC GetByIDStudent
// @Summary Get a Jurnal by Student Group ID
// @Description Get a Jurnal entry by Student Group ID
// @Produce json
// @Tags Jurnal
// @Param id path string true "Student Group ID"
// @Success 200 {object} education_management_service.GetJurnalResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals/student/{id} [get]
func (h *handler) GetJurnalByIDStudent(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.GroupId{Id: id}

	resp, err := h.grpcClient.JurnalService().GetByIDStudent(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListJurnal handles the RPC GetList
// @Summary Get a list of Jurnals
// @Description Get a list of Jurnal entries with pagination and search
// @Produce json
// @Tags Jurnal
// @Param limit query int false "Limit per page"
// @Param page query int false "Page number"
// @Param search query string false "Search query"
// @Success 200 {object} education_management_service.GetListJurnalResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals [get]
func (h *handler) GetListJurnal(c *gin.Context) {
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

	req := education_management_service.GetListJurnalRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.JurnalService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateJurnal handles the RPC UpdateJurnal
// @Summary Update a Jurnal by ID
// @Description Update a Jurnal entry by its ID
// @Accept json
// @Tags Jurnal
// @Produce json
// @Param id path string true "Jurnal ID"
// @Param request body education_management_service.UpdateJurnalRequest true "Jurnal update request"
// @Success 200 {object} education_management_service.GetJurnalResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals/{id} [put]
func (h *handler) UpdateJurnal(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateJurnalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.JurnalService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteJurnal handles the RPC Delete
// @Summary Delete a Jurnal by ID
// @Description Delete a Jurnal entry by its ID
// @Produce json
// @Tags Jurnal
// @Param id path string true "Jurnal ID"
// @Success 200 {object} education_management_service.JurnalEmpty
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
// @Router /api/v1/jurnals/{id} [delete]
func (h *handler) DeleteJurnal(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.JurnalID{Id: id}

	resp, err := h.grpcClient.JurnalService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
