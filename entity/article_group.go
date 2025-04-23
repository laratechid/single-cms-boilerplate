package entity

import "time"

type ArticleGroup struct {
	ID                   int32      `gorm:"primary_key" json:"id"`
	SequenceArticleGroup int        `json:"sequence_article_group"`
	HeadlineAt           *time.Time `json:"headline_at"`
	HeadlineBy           string     `json:"headline_by"`
	BreakingNewsAt       *time.Time `json:"breaking_news_at"`
	BreakingNewsBy       string     `json:"breaking_news_by"`
	GroupID              int32      `json:"group_id"`
	ArticleID            int64      `json:"article_id"`
	CreatedAt            *time.Time `json:"created_at"`
	CreatedBy            string     `json:"created_by"`
	UpdatedAt            *time.Time `json:"updated_at"`
	UpdatedBy            string     `json:"updated_by"`
	Group                Group      `json:"group"`
	Article              *Article   `json:"article,omitempty"`
}

func (t *ArticleGroup) TableName() string {
	return "article_groups"
}
