package controllers

import (
	"assigment2/server/models"
	"assigment2/server/repositories"
	"assigment2/server/views"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	repo repositories.OrdersRepo
}

func NewOrderController(repo repositories.OrdersRepo) *OrdersController {
	return &OrdersController{
		repo: repo,
	}
}

func WriteJsonResponse(ctx *gin.Context, payload *views.Response) {
	ctx.JSON(payload.Status, payload)
}

// Create Orders godoc
// @Summary    Create orders
// @Decription Create orders
// @Tags       orders
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllOrdersSwagger
// @Param Orders body models.ReqOrders true "Create Orders"
// @Router     /orders [post]
func (o *OrdersController) CreateOrders(ctx *gin.Context) {
	var req models.Orders
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errs := o.repo.CreateOrders(&req)
	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "CREATE_ORDERS_SUCCESS",
		Payload: req,
	})
}

// Get Orders godoc
// @Summary    Get orders
// @Decription Get orders
// @Tags       orders
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllOrdersSwagger
// @Router     /orders [get]
func (o *OrdersController) GetOrders(ctx *gin.Context) {
	orders, err := o.repo.GetOrders()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_ORDERS_FAIL",
			Error:   err.Error(),
		})
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_ORDERS_SUCCESS",
		Payload: orders,
	})
}

// Delete orders godoc
// @Summary    Delete orders
// @Decription Delete orders
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllOrdersSwagger
// @Param      id path int true "Orders ID"
// @Router     /orders/{id} [delete]
func (o *OrdersController) DeleteOrders(ctx *gin.Context) {
	getId := ctx.Params.ByName("id")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDERS_FAIL",
			Error:   errId.Error(),
		})
		return
	}
	errDel := o.repo.DeleteOrders(id)
	if errDel != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDERS_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_ORDERS_SUCCESS",
		Payload: nil,
	})
}

// Update Orders godoc
// @Summary    Update orders
// @Decription Update orders
// @Tags       orders
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllOrdersSwagger
// @Param      id path int true "Orders ID"
// @Param Orders body models.UpOrders true "Update Orders"
// @Router     /orders/{id} [patch]
func (o *OrdersController) UpdateOrders(ctx *gin.Context) {
	fmt.Print("UPDATE")
	getId := ctx.Params.ByName("id")
	var req models.Orders
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	errUpdate := o.repo.UpdateOrders(id, &req)
	if errUpdate != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PEOPLE_SUCCESS",
		Payload: req,
	})

}
