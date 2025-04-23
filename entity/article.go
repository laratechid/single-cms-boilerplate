package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	ID                        int64           `gorm:"primary_key;column:id" json:"id"`
	ArticleUUID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()" json:"article_uuid"`
	ArticleIDOld              int64           `json:"article_id_old"`
	DomainID                  int32           `json:"domain_id"`
	DomainIDOld               int32           `json:"domain_id_old"`
	FormatArticleID           int32           `json:"format_article_id"`
	SubRubricID               int32           `gorm:"column:sub_rubric_id" json:"sub_rubric_id"`
	Access                    string          `json:"access"`
	ContentCategory           string          `json:"content_category"`
	Language                  string          `json:"language"`
	TagTitle                  pq.StringArray  `gorm:"type:text[]" json:"tag_title"`
	TagArticle                pq.StringArray  `gorm:"type:text[]" json:"tag_article"`
	TitlePrint                string          `json:"title_print"`
	TitleDigital              string          `json:"title_digital"`
	TitleSEO                  string          `json:"title_seo"`
	Summary                   pq.StringArray  `gorm:"type:text[]" json:"summary"`
	Description               string          `json:"description"`
	Content                   string          `json:"content"`
	Upperdeck                 string          `json:"upperdeck"`
	Kicker                    string          `json:"kicker"`
	Taicing                   string          `json:"taicing"`
	IsByline                  bool            `json:"is_byline"`
	Snackbar                  string          `json:"snackbar"`
	IsActive                  bool            `json:"is_active"`
	WordCount                 int32           `json:"word_count"`
	AlphabetCount             int32           `json:"alphabet_count"`
	CanonicalURL              string          `json:"canonical_url"`
	CanonicalURLOld           string          `json:"canonical_url_old"`
	IsArticleJembatan         string          `json:"is_article_jembatan"`
	FootNote                  string          `json:"foot_note"`
	FeatureImage              string          `json:"feature_image"`
	FeatureImageCaption       string          `json:"feature_image_caption"`
	FactCheckClaim            string          `json:"fact_check_claim"`
	FactCheckResult           string          `json:"fact_check_result"`
	InfografisFile            string          `json:"infografis_file"`
	EmbedCode                 string          `json:"embed_code"`
	HeadlineAt                *time.Time      `json:"headline_at"`
	BreakingNewsAt            *time.Time      `json:"breaking_news_at"`
	PublishedAt               *time.Time      `json:"published_at"`
	PublishedBy               string          `json:"published_by"`
	ApprovedAt                *time.Time      `json:"approved_at"`
	ApprovedBy                string          `json:"approved_by"`
	Status                    string          `json:"status"`
	UnpublishedAt             *time.Time      `json:"unpublished_at"`
	UnpublishedBy             string          `json:"unpublished_by"`
	WritingStyle              string          `json:"writing_style"`
	CreatedAt                 *time.Time      `json:"created_at"`
	CreatedBy                 string          `json:"created_by"`
	UpdatedAt                 *time.Time      `json:"updated_at"`
	UpdatedBy                 string          `json:"updated_by"`
	DeletedAt                 *gorm.DeletedAt `json:"deleted_at"`
	DeletedBy                 string          `json:"deleted_by"`
	IsAnonymousReporterEditor bool            `json:"is_anonymous_reporter_editor"`
	TTSObjectPath             string          `json:"tts_object_path"`
	Groups                    *[]Group        `gorm:"many2many:article_groups;foreignKey:id;joinForeignKey:article_id;references:id;joinReferences:group_id" json:"groups,omitempty"`
	Attachments               *[]Attachment   `gorm:"many2many:articles_attachments;foreignKey:id;joinForeignKey:article_id;references:id;joinReferences:attachment_id"`
	ArticleUser               *[]ArticleUser  `gorm:"foreignKey:ArticleID" json:"article_user,omitempty"`
	SubRubric                 SubRubric       `gorm:"ForeignKet:SubRubricID;references:id" json:"sub_rubric,omitempty"`
	ArticleGroups             *[]ArticleGroup `gorm:"foreignKey:ArticleID" json:"article_groups,omitempty"`
}
