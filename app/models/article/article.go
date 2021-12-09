package article

import (
	"blog/app/models"
	"blog/pkg/model"
	"gorm.io/gorm"
)

type Article struct {
	models.BaseModel
	Name   string `gorm:"column:name;type:varchar(200);not null;"`
	Author string `gorm:"column:author;type:varchar(30);not null;default:Aron"`
	Cid    uint64 `gorm:"column:cid;not null"`
	Like   uint64 `gorm:"column:like;not null"`
	View   uint64 `gorm:"column:view;type:uint;not null;default:0"`
	Status uint8  `gorm:"column:status;type:uint;default:1;not null"`
	Sort   uint32 `gorm:"column:sort;type:uint;default:0;not null""`
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func GetByCid(id int64) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("cid=?", id).Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func GetById(id int64) (Article, error) {
	var article Article
	if err := model.DB.Where("id=?", id).Find(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}

func UpdateReview(id int64) bool {
	article, err := GetById(id)
	if err != nil {
		return false
	}
	db := model.DB.Model(&article).Where("id=?", id).
		UpdateColumn("view", gorm.Expr("view + ?", 1))
	return db.RowsAffected != 0
}
