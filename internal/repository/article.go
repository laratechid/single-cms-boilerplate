package repository

import (
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetByID(id int64) (entity.Article, error)
	Create(entity entity.Article) error
	Update(entity entity.Article) error
	GetAll(p dto.PaginationRequestDto) ([]entity.Article, int64, error)
	Delete(id int64) error
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

func (r articleRepository) Create(entity entity.Article) error {
	create := r.db.Create(&entity)
	err := create.Error
	if err != nil {
		helper.LogErr(err)
		return err
	}
	return nil
}

func (r articleRepository) Update(entity entity.Article) error {
	update := r.db.Updates(&entity)
	err := update.Error
	if err != nil {
		helper.LogErr(err)
		return err
	}
	return nil
}

func (r articleRepository) GetAll(p dto.PaginationRequestDto) ([]entity.Article, int64, error) {
	article := []entity.Article{}
	offset := (p.Page - 1) * p.Limit
	query := r.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Offset(offset).Limit(p.Limit)
	}).Find(&article)
	var total int64
	count := query.Count(&total)
	err := count.Error
	if err != nil {
		helper.LogErr(err)
		return nil, total, err
	}
	err = query.Error
	if err != nil {
		helper.LogErr(err)
		return nil, total, err
	}
	return article, total, nil
}

func (r articleRepository) Delete(id int64) error {
	article := entity.Article{ID: id}
	err := r.db.Delete(&article).Error
	if err != nil {
		helper.LogErr(err)
		return err
	}
	return nil
}
