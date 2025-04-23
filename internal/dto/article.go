package dto

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID                        int64             `json:"id"`
	ArticleUUID               uuid.UUID         `json:"article_uuid"`
	ArticleIDOld              int64             `json:"article_id_old"`
	DomainID                  int32             `json:"domain_id"`
	DomainIDOld               int32             `json:"domain_id_old"`
	FormatArticleID           int32             `json:"format_article_id"`
	SubRubricID               int32             `json:"sub_rubric_id"`
	Access                    string            `json:"access"`
	ContentCategory           string            `json:"content_category"`
	Language                  string            `json:"language"`
	TagTitle                  []string          `json:"tag_title"`
	TagArticleNew             []string          `json:"tag_article_new"`
	TagArticle                []*ArticleTagsOld `json:"tag_article"`
	TitlePrint                string            `json:"title_print"`
	TitleDigital              string            `json:"title_digital"`
	TitleSEO                  string            `json:"title_seo"`
	Summary                   []string          `json:"summary"`
	Description               string            `json:"description"`
	Content                   string            `json:"content"`
	Upperdeck                 string            `json:"upperdeck"`
	Kicker                    string            `json:"kicker"`
	Taicing                   string            `json:"taicing"`
	IsByline                  bool              `json:"is_byline"`
	SnackBar                  string            `json:"snackbar"`
	IsActive                  bool              `json:"is_active"`
	WordCount                 int32             `json:"word_count"`
	AlphabetCount             int32             `json:"alphabet_count"`
	CanonicalURL              string            `json:"canonical_url"`
	CanonicalURLOld           string            `json:"canonical_url_old"`
	IsArticleJembatan         string            `json:"is_article_jembatan"`
	FootNote                  string            `json:"foot_note"`
	FeatureImage              string            `json:"feature_image"`
	FeatureImageCaption       string            `json:"feature_image_caption"`
	FactCheckClaim            string            `json:"fact_check_claim"`
	FactCheckResult           string            `json:"fact_check_result"`
	FactCheckResultImage      string            `json:"fact_check_result_image"`
	InfografisFile            string            `json:"infografis_file"`
	EmbedCode                 string            `json:"embed_code"`
	HeadlineAt                string            `json:"headline_at"`
	BreakingNewsAt            string            `json:"breaking_news_at"`
	PublishedAt               *time.Time        `json:"published_at"`
	PublishedBy               string            `json:"published_by"`
	ApprovedAt                string            `json:"approved_at"`
	ApprovedBy                string            `json:"approved_by"`
	Status                    string            `json:"status"`
	UnpublishedAt             string            `json:"unpublished_at"`
	UnpublishedBy             string            `json:"unpublished_by"`
	WritingStyle              string            `json:"writing_style"`
	CreatedAt                 string            `json:"created_at,omitempty"`
	CreatedBy                 string            `json:"created_by,omitempty"`
	UpdatedAt                 string            `json:"updated_at,omitempty"`
	UpdatedBy                 string            `json:"updated_by,omitempty"`
	IsLabelTag                string            `json:"is_label_tag"` // will be deleted soon
	IsAnonymousReporterEditor bool              `json:"is_anonymous_reporter_editor"`
	IsAudioAvailable          bool              `json:"is_audio_available"`
	TTSAudioURL               string            `json:"tts_audio_url"`
	Paragraphs                []string          `json:"paragraphs"`
	Attachments               []*Attachment     `json:"attachments"`
	ArticleUser               []*ArticleUser    `json:"article_user"`
	SubRubric                 *SubRubric        `json:"sub_rubric"`
	ArticleGroups             []*ArticleGroup   `json:"article_groups"`
}

type ArticleGroup struct {
	ID                   int32      `json:"id"`
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

type ArticleUser struct {
	ID        int32      `json:"id"`
	ArticleID int32      `json:"article_id"`
	UserID    int32      `json:"user_id"`
	Type      string     `json:"type"`
	CreatedAt *time.Time `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy string     `json:"deleted_by"`
	User      *User      `json:"user"`
	Article   *Article   `json:"article"`
}
