// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameMonitor = "monitors"

// Monitor mapped from table <monitors>
type Monitor struct {
	MonitorID         int64           `gorm:"column:monitor_id;primaryKey" json:"monitor_id"`
	MonitorProducerID int64           `gorm:"column:monitor_producer_id;not null" json:"monitor_producer_id"`
	Model             string          `gorm:"column:model;not null" json:"model"`
	MonitorProducer   MonitorProducer `json:"monitor_producer"`
	PcTypes           []PcType        `json:"pc_types"`
}

// TableName Monitor's table name
func (*Monitor) TableName() string {
	return TableNameMonitor
}
