package usecase

import (
	"context"

	"github.com/madmuzz05/golang-assigment-2/internal/database/gorm/postgres"
	sysresponse "github.com/madmuzz05/golang-assigment-2/pkg/helper/sys_response"
	"github.com/madmuzz05/golang-assigment-2/service/dto/request"
	"github.com/madmuzz05/golang-assigment-2/service/dto/response"
	"github.com/madmuzz05/golang-assigment-2/service/repository"
)

type OrderUsecase struct {
	OrderRepository repository.IOrderRepository
	GormDB          *postgres.GormDB
}

func InitOrderUsecase(OrderRepository repository.IOrderRepository, gormDb *postgres.GormDB) IOrderUsecase {
	return &OrderUsecase{
		OrderRepository: OrderRepository,
		GormDB:          gormDb,
	}
}

type IOrderUsecase interface {
	GetOrders(ctx context.Context) (res []response.OrdersDtoResponse, err sysresponse.IError)
	StoreOrder(ctx context.Context, req request.OrderDtoRequest, id string) (res response.OrdersDtoResponse, err sysresponse.IError)
	GetOrder(ctx context.Context, id string) (res response.OrdersDtoResponse, err sysresponse.IError)
	DeleteOrder(ctx context.Context, id string) (err sysresponse.IError)
}
