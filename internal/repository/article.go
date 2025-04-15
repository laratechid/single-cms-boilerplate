package repository

import (
	"super-cms/helper"
	"super-cms/internal/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetByID(id int64) (entity.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r articleRepository) GetByID(id int64) (entity.Article, error) {
	article := entity.Article{ID: id}
	err := r.db.First(&article).Error
	if err != nil {
		helper.LogErr(err)
	}
	return article, err
}
