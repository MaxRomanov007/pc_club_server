// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNamePcType = "pc_types"

// PcType mapped from table <pc_types>
type PcType struct {
	PcTypeID     int64         `gorm:"column:pc_type_id;primaryKey" json:"pc_type_id"`
	ProcessorID  int64         `gorm:"column:processor_id;not null" json:"processor_id"`
	VideoCardID  int64         `gorm:"column:video_card_id;not null" json:"video_card_id"`
	MonitorID    int64         `gorm:"column:monitor_id;not null" json:"monitor_id"`
	RAMID        int64         `gorm:"column:ram_id;not null" json:"ram_id"`
	Name         string        `gorm:"column:name;not null" json:"name"`
	Description  string        `gorm:"column:description" json:"description"`
	HourCost     float32       `gorm:"column:hour_cost;not null" json:"hour_cost"`
	Processor    Processor     `json:"processor"`
	VideoCard    VideoCard     `json:"video_card"`
	Monitor      Monitor       `json:"monitor"`
	RAM          RAM           `json:"ram"`
	Pcs          []Pc          `json:"pcs"`
	PcTypeImages []PcTypeImage `json:"pc_type_images"`
}

// TableName PcType's table name
func (*PcType) TableName() string {
	return TableNamePcType
}
