package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api_gateway/genproto/education_management_service"
)

// @Security ApiKeyAuth
// @Router /student_payment/{id} [get]
// @Summary Get student payment by ID
// @Description API for getting a student payment by ID
// @Tags student_payment
// @Accept json
// @Produce json
// @Param id path string true "Student Payment ID"
// @Success 200 {object} education_management_service.GetStudentPaymentResponse
// @Failure 404 {object} models.ResponseError "Student Payment not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetStudentPaymentByID(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.StudentPaymentID{Id: id}

	resp, err := h.grpcClient.StudentPaymentService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student_payment/list [get]
// @Summary Get list of student payments
// @Description API for getting a list of student payments
// @Tags student_payment
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} education_management_service.GetListStudentPaymentResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListStudentPayment(c *gin.Context) {
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

	req := education_management_service.GetListStudentPaymentRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.StudentPaymentService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /create/student_payment [post]
// @Summary Create student payment
// @Description API for creating a student payment
// @Tags student_payment
// @Accept json
// @Produce json
// @Param student_payment body education_management_service.CreateStudentPaymentRequest true "Student Payment"
// @Success 200 {object} education_management_service.StudentPaymentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateStudentPayment(c *gin.Context) {
	var req education_management_service.CreateStudentPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.StudentPaymentService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student_payment/{id} [put]
// @Summary Update student payment by ID
// @Description API for updating a student payment by ID
// @Tags student_payment
// @Accept json
// @Produce json
// @Param id path string true "Student Payment ID"
// @Param student_payment body education_management_service.UpdateStudentPaymentRequest true "Student Payment"
// @Success 200 {object} education_management_service.GetStudentPaymentResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Student Payment not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateStudentPayment(c *gin.Context) {
	id := c.Param("id")
	var req education_management_service.UpdateStudentPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.StudentPaymentService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /student_payment/{id} [delete]
// @Summary Delete student payment by ID
// @Description API for deleting a student payment by ID
// @Tags student_payment
// @Accept json
// @Produce json
// @Param id path string true "Student Payment ID"
// @Success 200 {object} education_management_service.StudentPaymentEmpty
// @Failure 404 {object} models.ResponseError "Student Payment not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteStudentPayment(c *gin.Context) {
	id := c.Param("id")
	req := &education_management_service.StudentPaymentID{Id: id}

	resp, err := h.grpcClient.StudentPaymentService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
