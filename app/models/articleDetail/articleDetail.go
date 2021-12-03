package articleDetail

import (
	"blog/pkg/model"
	"fmt"
)

type ArticleDetail struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`
	ArticleId uint64 `gorm:"column:article_id;not null"`
	Content   string `gorm:"type:text;column:content;not null"`
}

func GetById(id int64) (ArticleDetail, error) {
	var articleDetail ArticleDetail
	if err := model.DB.Where("article_id=?", id).First(&articleDetail).Error; err != nil {
		fmt.Println(err.Error())
		return articleDetail, err
	}
	return articleDetail, nil
}
