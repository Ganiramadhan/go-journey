package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string         `gorm:"type:char(36);primaryKey" json:"id"`
	Username      string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password      string         `gorm:"type:varchar(255);not null" json:"-"`
	FullName      string         `gorm:"type:varchar(150);not null" json:"full_name"`
	Role          string         `gorm:"type:varchar(20);default:'guest';not null" json:"role"`
	RegisterDate  time.Time      `gorm:"autoCreateTime" json:"register_date"`
	EsignID       string         `gorm:"type:varchar(100)" json:"esign_id"`
	EsignStatusID string         `gorm:"type:varchar(50)" json:"esign_status_id"`
	RefreshToken  string         `gorm:"type:varchar(255)" json:"-"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func (User) TableName() string {
	return "users"
}
