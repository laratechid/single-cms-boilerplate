package service

import (
	"super-cms/internal/dto"
	"super-cms/internal/entity"
	"super-cms/internal/repository"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestArticleService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepo := repository.NewMockArticleRepository(ctrl)

	articleService := &articleService{
		articleRepo: mockArticleRepo,
	}

	t.Run("Article found", func(t *testing.T) {
		article := entity.Article{ID: 1, Title: "Test Article", Body: "Test Content"}
		articleResponse := dto.ArticleDetailResponse{}
		copier.Copy(&articleResponse, &article)
		mockArticleRepo.EXPECT().GetByID(int64(article.ID)).Return(article, nil)
		result, err := articleService.GetByID(int64(article.ID))
		assert.NoError(t, err)
		assert.Equal(t, &articleResponse, result)
	})
}
