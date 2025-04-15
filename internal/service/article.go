package service

import (
	"super-cms/helper"
	"super-cms/internal/dto"
	"super-cms/internal/repository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ArticleService interface {
	GetByID(id int64) (*dto.ArticleDetailResponse, error)
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{
		articleRepo: repository.NewArticleRepository(db),
	}
}

func (s *articleService) GetByID(id int64) (*dto.ArticleDetailResponse, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		helper.LogErr(err)
		return nil, err
	}
	response := dto.ArticleDetailResponse{}
	err = copier.Copy(&response, &article)
	if err != nil {
		helper.LogErr(err)
		return nil, err
	}
	return &response, nil
}
