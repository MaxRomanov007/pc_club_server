// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNamePcOrder = "pc_orders"

// PcOrder mapped from table <pc_orders>
type PcOrder struct {
	PcOrderID       int64         `gorm:"column:pc_order_id;primaryKey" json:"pc_order_id"`
	UserID          int64         `gorm:"column:user_id;not null" json:"user_id"`
	PcID            int64         `gorm:"column:pc_id;not null" json:"pc_id"`
	PcOrderStatusID int64         `gorm:"column:pc_order_status_id;not null" json:"pc_order_status_id"`
	Code            string        `gorm:"column:code;not null" json:"code"`
	Cost            float32       `gorm:"column:cost;not null" json:"cost"`
	StartTime       time.Time     `gorm:"column:start_time;not null" json:"start_time"`
	EndTime         time.Time     `gorm:"column:end_time;not null" json:"end_time"`
	Duration        int16         `gorm:"column:duration;not null" json:"duration"`
	ActualEndTime   time.Time     `gorm:"column:actual_end_time;not null" json:"actual_end_time"`
	OrderDate       time.Time     `gorm:"column:order_date;not null;default:getdate()" json:"order_date"`
	User            User          `json:"user"`
	Pc              Pc            `json:"pc"`
	PcOrderStatus   PcOrderStatus `json:"pc_order_status"`
}

// TableName PcOrder's table name
func (*PcOrder) TableName() string {
	return TableNamePcOrder
}
