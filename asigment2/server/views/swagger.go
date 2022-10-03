package views

import (
	"assigment2/server/models"
)

type GetAllOrdersSwagger struct {
	Response
	Payload []models.ReqOrders
}
