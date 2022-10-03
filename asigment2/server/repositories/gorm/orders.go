package gorm

import (
	"assigment2/server/models"
	"assigment2/server/repositories"
	"fmt"

	"gorm.io/gorm"
)

type ordersRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrdersRepo {
	return &ordersRepo{db: db}
}

func (o *ordersRepo) CreateOrders(orders *models.Orders) error {
	return o.db.Create(orders).Error
}

func (o *ordersRepo) GetOrders() (*[]models.Orders, error) {
	var orders []models.Orders
	err := o.db.Model(&models.Orders{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (o *ordersRepo) DeleteOrders(id int) error {
	orders := models.Orders{}
	items := models.Items{}
	errO := o.db.Where("id = ?", id).Delete(&orders).Error
	if errO != nil {
		return errO
	}

	errI := o.db.Where("orderid = ?", id).Delete(&items).Error
	if errO != nil {
		return errI
	}

	return nil
}

func (o *ordersRepo) UpdateOrders(id int, orders *models.Orders) error {
	order := models.Orders{}
	err := o.db.Model(&order).Where("id = ?", id).Updates(models.Orders{CustomerName: orders.CustomerName}).Error
	if err != nil {
		return err
	}
	item := models.Items{}
	for _, dataItem := range orders.Items {
		fmt.Println("--------------------")
		fmt.Println(dataItem)
		errItm := o.db.Model(&item).Where("id = ?", dataItem.ID).Updates(models.Items{Item_code: dataItem.Item_code, Description: dataItem.Description, Quantity: dataItem.Quantity}).Error
		if err != nil {
			return errItm
		}
	}
	return nil
}
