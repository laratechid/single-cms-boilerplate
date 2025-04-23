package dto

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 int64      `json:"id"`
	UUID               uuid.UUID  `json:"uuid"`
	UserRempID         int32      `json:"user_remp_id"`
	IsEditor           bool       `json:"is_editor"`
	IsContributor      bool       `json:"is_contributor"`
	IsAdmin            bool       `json:"is_admin"`
	Nik                string     `json:"nik"`
	Name               string     `json:"name"`
	Alias              string     `json:"alias"`
	Foto               string     `json:"foto"`
	Biodata            string     `json:"biodata"`
	BiodataEn          string     `json:"biodata_en"`
	Email              string     `json:"email"`
	MainEmail          string     `json:"main_email"`
	SecondaryEmail     string     `json:"secondary_email"`
	Username           string     `json:"username"`
	Password           string     `json:"password"`
	ForceResetPassword bool       `json:"force_reset_password"`
	Facebook           string     `json:"facebook"`
	Twitter            string     `json:"twitter"`
	Linked             string     `json:"linked"`
	ActivedAt          *time.Time `json:"actived_at"`
	ActivedBy          string     `json:"actived_by"`
	CreatedAt          *time.Time `json:"created_at"`
	CreatedBy          string     `json:"created_by"`
	UpdatedAt          *time.Time `json:"updated_at"`
	UpdatedBy          string     `json:"updated_by"`
	DeletedAt          *time.Time `json:"deleted_at"`
	DeletedBy          string     `json:"deleted_by"`
}
