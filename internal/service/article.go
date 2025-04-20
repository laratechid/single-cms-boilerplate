package service

import (
	"math"
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/entity"
	"super-cms/internal/repository"

	"github.com/go-stack/stack"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll(p dto.PaginationRequestDto) (*dto.PaginationResponseDto[dto.ArticleDetailResponse], error)
	GetByID(id int64) (*dto.ArticleDetailResponse, error)
	Create(dto dto.ArticleCreateRequestDto) error
	Update(id int64, dto dto.ArticleUpdateRequestDto) error
	Delete(id int64) error
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{
		articleRepo: repository.NewArticleRepository(db),
	}
}

func (r *articleService) traceErr(err error) {
	stack := stack.Caller(1).Frame().Function
	helper.LogErr(err, stack)
}

func (s *articleService) GetAll(p dto.PaginationRequestDto) (*dto.PaginationResponseDto[dto.ArticleDetailResponse], error) {
	data, total, err := s.articleRepo.GetAll(p)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	articles := []dto.ArticleDetailResponse{}
	if err = copier.Copy(&articles, &data); err != nil {
		s.traceErr(err)
		return nil, err
	}
	response := dto.PaginationResponseDto[dto.ArticleDetailResponse]{
		List:          articles,
		Limit:         int64(p.Limit),
		TotalEntry:    total,
		TotalPage:     int64(math.Ceil(float64(total) / float64(p.Limit))),
		IsHasNextPage: total > (int64(p.Limit) * int64(p.Page)),
	}
	return &response, nil
}

func (s *articleService) GetByID(id int64) (*dto.ArticleDetailResponse, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		s.traceErr(err)
		return nil, err
	}
	response := dto.ArticleDetailResponse{}
	if err = copier.Copy(&response, &article); err != nil {
		s.traceErr(err)
		return nil, err
	}
	return &response, nil
}

func (s *articleService) Create(dto dto.ArticleCreateRequestDto) error {
	entity := entity.Article{}
	if err := copier.Copy(&entity, &dto); err != nil {
		s.traceErr(err)
		return err
	}
	if err := s.articleRepo.Create(entity); err != nil {
		s.traceErr(err)
		return err
	}
	return nil
}

func (s *articleService) Update(id int64, dto dto.ArticleUpdateRequestDto) error {
	entity := entity.Article{ID: id}
	if err := copier.Copy(&entity, &dto); err != nil {
		s.traceErr(err)
		return err
	}
	if err := s.articleRepo.Update(entity); err != nil {
		s.traceErr(err)
		return err
	}
	return nil
}

func (s *articleService) Delete(id int64) error {
	err := s.articleRepo.Delete(id)
	if err != nil {
		s.traceErr(err)
		return err
	}
	return nil
}
