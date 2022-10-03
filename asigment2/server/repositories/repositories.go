package repositories

import "assigment2/server/models"

type OrdersRepo interface {
	CreateOrders(orders *models.Orders) error
	GetOrders() (*[]models.Orders, error)
	DeleteOrders(id int) error
	UpdateOrders(id int, orders *models.Orders) error
}
