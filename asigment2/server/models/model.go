package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	CustomerName string
	OrderedAt    time.Time
	Items        []Items `gorm:"foreignKey:Orderid"`
}

type Items struct {
	gorm.Model
	Item_code   string
	Description string
	Quantity    int
	Orderid     int
}

type ReqOrders struct {
	CustomerName string
	Items        []ReqItems
}

type ReqItems struct {
	Item_code   string
	Description string
	Quantity    int
}

type UpOrders struct {
	CustomerName string
	Items        []UpItems
}

type UpItems struct {
	Item_code   string
	Description string
	Quantity    int
	Id          int
}
