// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameDishOrderStatus = "dish_order_statuses"

// DishOrderStatus mapped from table <dish_order_statuses>
type DishOrderStatus struct {
	DishOrderStatusID int64       `gorm:"column:dish_order_status_id;primaryKey" json:"dish_order_status_id"`
	Name              string      `gorm:"column:name;not null" json:"name"`
	DishOrders        []DishOrder `json:"dish_orders"`
}

// TableName DishOrderStatus's table name
func (*DishOrderStatus) TableName() string {
	return TableNameDishOrderStatus
}
