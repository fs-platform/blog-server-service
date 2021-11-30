package category

import (
	"blog/app/models"
	"blog/pkg/model"
)

type Categories struct {
	models.BaseModel
	Name   string `gorm:"column:name;type:varchar(10);not null"`
	Status uint8  `gorm:"column:status;type:uint;default:1;not null"`
	Sort   uint32 `gorm:"column:sort;type:uint;default:0;not null""`
}

func GetAll() ([]Categories, error) {
	var categories []Categories
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}
