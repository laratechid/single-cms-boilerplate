package repository

import (
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/entity"

	"github.com/go-stack/stack"
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
	return &articleRepository{
		db,
	}
}

func (r articleRepository) stack() string {
	return stack.Caller(0).Frame().Function
}

func (r articleRepository) GetByID(id int64) (entity.Article, error) {
	article := entity.Article{ID: id}
	err := r.db.First(&article).Error
	if err != nil {
		helper.LogErr(err, r.stack())
	}
	return article, err
}

func (r articleRepository) Create(entity entity.Article) error {
	create := r.db.Create(&entity)
	err := create.Error
	if err != nil {
		helper.LogErr(err, r.stack())
		return err
	}
	return nil
}

func (r articleRepository) Update(entity entity.Article) error {
	update := r.db.Updates(&entity)
	err := update.Error
	if err != nil {
		helper.LogErr(err, r.stack())
		return err
	}
	return nil
}

func (r articleRepository) GetAll(p dto.PaginationRequestDto) ([]entity.Article, int64, error) {
	var total int64
	article := []entity.Article{}
	p.SetDefault()
	query := r.db.Offset(p.Offset).Limit(p.Limit).Find(&article)
	r.db.Model(&entity.Article{}).Count(&total)
	err := query.Error
	if err != nil {
		helper.LogErr(err, r.stack())
		return nil, total, err
	}
	return article, total, nil
}

func (r articleRepository) Delete(id int64) error {
	article := entity.Article{ID: id}
	err := r.db.Delete(&article).Error
	if err != nil {
		helper.LogErr(err, r.stack())
		return err
	}
	return nil
}
