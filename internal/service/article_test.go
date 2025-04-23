package service

import (
	"errors"
	"testing"

	"go-pustaka-api/entity"
	"go-pustaka-api/internal/dto"
	"go-pustaka-api/internal/repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestArticleService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockArticleRepository(ctrl)
	articleService := &articleService{articleRepo: mockRepo}

	t.Run("success", func(t *testing.T) {
		pagination := dto.PaginationRequestDto{
			Page:  1,
			Limit: 10,
		}

		articles := []entity.Article{
			{ID: 1},
			{ID: 2},
		}
		total := int64(2)

		mockRepo.EXPECT().GetAll(pagination).Return(articles, total, nil)

		result, err := articleService.GetAll(pagination)

		assert.NoError(t, err)
		assert.Equal(t, total, result.TotalEntry)
		assert.Len(t, result.List, 2)
	})

	t.Run("repository error", func(t *testing.T) {
		pagination := dto.PaginationRequestDto{
			Page:  1,
			Limit: 10,
		}

		mockRepo.EXPECT().GetAll(pagination).Return(nil, int64(0), errors.New("repository error"))

		_, err := articleService.GetAll(pagination)

		assert.Error(t, err)
	})
}

func TestArticleService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockArticleRepository(ctrl)
	articleService := &articleService{articleRepo: mockRepo}

	t.Run("success", func(t *testing.T) {
		id := int64(1)
		article := entity.Article{
			ID: id,
		}

		mockRepo.EXPECT().GetByID(id).Return(article, nil)

		result, err := articleService.GetByID(id)

		assert.NoError(t, err)
		assert.Equal(t, id, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		id := int64(999)
		mockRepo.EXPECT().GetByID(id).Return(entity.Article{}, gorm.ErrRecordNotFound)

		_, err := articleService.GetByID(id)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})
}

func TestArticleService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockArticleRepository(ctrl)
	articleService := &articleService{articleRepo: mockRepo}

	t.Run("success", func(t *testing.T) {
		id := int64(1)
		mockRepo.EXPECT().Delete(id).Return(nil)

		err := articleService.Delete(id)

		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		id := int64(999)
		mockRepo.EXPECT().Delete(id).Return(gorm.ErrRecordNotFound)

		err := articleService.Delete(id)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})
}
