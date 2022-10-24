package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Asliddin3/exam-api-gateway/genproto/customer"
	l "github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary create customer with info
// @Description this func create customer with
// @Tags customer
// @Accept json
// @Produce json
// @Param customer body customer.CustomerRequest true "Customer"
// @Success 200 {object} customer.CustomerResponse
// @Router /customer [post]
func (h *handlerV1) CreateCustomer(c *gin.Context) {
	var (
		body        customer.CustomerRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().CreateCustomer(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create customer", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary update customers
// @Description this func update customers
// @Tags customer
// @Accept json
// @Produce json
// @Param customer body customer.CustomerResponse true "Customer"
// @Success 200 {object} customer.CustomerResponse
// @Router /customer/update [patch]
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var (
		body        customer.CustomerResponse
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().UpdateCustomer(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update customer", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary delete customer
// @Description this func delete customer
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /customer/delete/{id} [delete]
func (h *handlerV1) DeleteCustomer(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &customer.CustomerId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().DeleteCustomer(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete customer", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary get all customers
// @Description this func get all customers
// @Tags customer
// @Accept json
// @Produce json
// @Success 200 {object} customer.ListCustomers
// @Router /customer/list [get]
func (h *handlerV1) GetListCustomers(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().GetListCustomers(ctx, &customer.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list customers", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary get customer with post
// @Description this func get customer with post
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} customer.CustomerResponsePost
// @Router /customer/post/{id} [get]
func (h *handlerV1) GetCustomerPostById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &customer.CustomerId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().GetById(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get customer post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1

// @Summary get customer info
// @Description this func get customer info
// @Tags customer
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} customer.CustomerResponse
// @Router /customer/{id} [get]
func (h *handlerV1) GetCustomerInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &customer.CustomerId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.CustomerService().GetCustomerInfo(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get customer info", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
