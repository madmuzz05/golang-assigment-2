package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/golang-assigment-2/pkg/helper/sys_response"
	"github.com/madmuzz05/golang-assigment-2/service/dto/request"
)

func (h OrderHandler) GetOrders(ctx *gin.Context) {

	res, err := h.OrderUsecase.GetOrders(ctx)
	if err != nil {
		sysresponse.GetSuccessMessage(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetSuccessMessage(ctx, http.StatusOK, "Found Data", res)
}

func (h OrderHandler) GetOrder(ctx *gin.Context) {

	id := ctx.Param("order_id")
	if id == "" {
		sysresponse.GetSuccessMessage(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "ID is required")
	}

	res, err := h.OrderUsecase.GetOrder(ctx, id)
	if err != nil {
		sysresponse.GetSuccessMessage(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetSuccessMessage(ctx, http.StatusOK, "Found Data", res)
}
func (h OrderHandler) DeleteOrder(ctx *gin.Context) {

	id := ctx.Param("order_id")
	if id == "" {
		sysresponse.GetSuccessMessage(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "ID is required")
	}

	err := h.OrderUsecase.DeleteOrder(ctx, id)
	if err != nil {
		sysresponse.GetSuccessMessage(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetSuccessMessage(ctx, http.StatusOK, "Delete data success", nil)
}
func (h OrderHandler) StoreOrder(ctx *gin.Context) {
	id := ctx.Param("order_id")
	req := request.OrderDtoRequest{}
	if valErr := ctx.BindJSON(&req); valErr != nil {
		sysresponse.GetSuccessMessage(ctx, http.StatusServiceUnavailable, "data tidak sesuai", valErr.Error())
		return
	}
	if id == "" && req.OrderedAt == "" {
		sysresponse.GetSuccessMessage(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "OrderedAt is required")
		return
	}

	if len(req.Items) == 0 {
		sysresponse.GetSuccessMessage(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "Item is required")
		return
	}

	res, err := h.OrderUsecase.StoreOrder(ctx, req, id)
	if err != nil {
		sysresponse.GetSuccessMessage(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}
	if id != "" {
		sysresponse.GetSuccessMessage(ctx, http.StatusOK, "Update data success", res)
		return
	}
	sysresponse.GetSuccessMessage(ctx, http.StatusOK, "Create data success", nil)
}
