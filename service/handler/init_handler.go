package handler

import "github.com/madmuzz05/golang-assigment-2/service/usecase"

type OrderHandler struct {
	OrderUsecase usecase.IOrderUsecase
}

func InitOrderHandler(orderHandler usecase.IOrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: orderHandler,
	}
}
