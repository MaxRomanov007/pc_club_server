// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameDishImage = "dish_images"

// DishImage mapped from table <dish_images>
type DishImage struct {
	DishImageID int64  `gorm:"column:dish_image_id;primaryKey" json:"dish_image_id"`
	DishID      int64  `gorm:"column:dish_id;not null" json:"dish_id"`
	IsMain      bool   `gorm:"column:is_main;not null;default:0" json:"is_main"`
	Path        string `gorm:"column:path" json:"path"`
	Dish        Dish   `json:"dish"`
}

// TableName DishImage's table name
func (*DishImage) TableName() string {
	return TableNameDishImage
}
