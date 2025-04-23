package repository

import (
	"super-cms/entity"
	"super-cms/helper"
	"super-cms/internal/dto"

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
	return &articleRepository{db}
}

func (r articleRepository) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (r articleRepository) GetByID(id int64) (entity.Article, error) {
	query := r.db.Where(entity.Article{ID: id})
	article := entity.Article{}
	err := query.Preload("Attachments", func(query *gorm.DB) *gorm.DB {
		query.Select("article_attachments.id", "article_attachments.article_id", "article_attachments.attachment_id")
		return query
	}).
		Preload("SubRubric", func(query *gorm.DB) *gorm.DB {
			query = query.Select([]string{
				"sub_rubrics.id",
				"sub_rubrics.name",
				"sub_rubrics.alias",
				"sub_rubrics.rubric_id",
			}).
				Preload("Rubric", func(query *gorm.DB) *gorm.DB {
					query = query.Select("rubrics.id", "rubrics.name", "rubrics.alias")
					return query
				})
			return query
		}).
		Preload("ArticleGroups", func(query *gorm.DB) *gorm.DB {
			query = query.Select("article_groups.id", "article_groups.id", "article_groups.group_id", "article_groups.article_id").
				Preload("Group", func(query *gorm.DB) *gorm.DB {
					query = query.Select([]string{
						"groups.id",
						"groups.uuid",
						"groups.cover",
						"groups.caption",
						"groups.alias",
						"groups.description",
						"groups.category",
						"groups.sequence",
						"groups.published_at",
						"groups.created_at",
						"groups.status",
						"groups.date",
					})
					return query
				})
			return query
		}).
		Preload("ArticleUser", func(query *gorm.DB) *gorm.DB {
			query = query.Select("article_user.id", "article_user.article_id", "article_user.user_id", "article_user.type").
				Preload("User", func(query *gorm.DB) *gorm.DB {
					query = query.Select([]string{
						"users.id",
						"users.uuid",
						"users.name",
						"users.alias",
						"users.foto",
						"users.username",
						"users.email",
						"users.nik",
						"users.biodata",
					})
					return query
				})
			return query
		}).
		Where("domain_id = ?", 1).
		// Scopes(domain.PublishedCondition(c)). // check platform "cms"
		First(&article).Error

	if err != nil {
		r.traceErr(err)
		return article, err
	}

	return article, nil
}

func (r articleRepository) Create(entity entity.Article) error {
	create := r.db.Create(&entity)
	err := create.Error
	if err != nil {
		r.traceErr(err)
		return err
	}
	return nil
}

func (r articleRepository) Update(entity entity.Article) error {
	update := r.db.Updates(&entity)
	err := update.Error
	if err != nil {
		r.traceErr(err)
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
		r.traceErr(err)
		return nil, total, err
	}
	return article, total, nil
}

func (r articleRepository) Delete(id int64) error {
	article := entity.Article{ID: id}
	err := r.db.Delete(&article).Error
	if err != nil {
		r.traceErr(err)
		return err
	}
	return nil
}
