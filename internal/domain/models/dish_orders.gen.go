// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameDishOrder = "dish_orders"

// DishOrder mapped from table <dish_orders>
type DishOrder struct {
	DishOrderID       int64           `gorm:"column:dish_order_id;primaryKey" json:"dish_order_id"`
	DishOrderStatusID int64           `gorm:"column:dish_order_status_id;not null" json:"dish_order_status_id"`
	UserID            int64           `gorm:"column:user_id;not null" json:"user_id"`
	Cost              float32         `gorm:"column:cost;not null" json:"cost"`
	OrderDate         time.Time       `gorm:"column:order_date;not null;default:getdate()" json:"order_date"`
	DishOrderStatus   DishOrderStatus `json:"dish_order_status"`
	User              User            `json:"user"`
	DishOrderList     []DishOrderList `json:"dish_order_list"`
}

// TableName DishOrder's table name
func (*DishOrder) TableName() string {
	return TableNameDishOrder
}
