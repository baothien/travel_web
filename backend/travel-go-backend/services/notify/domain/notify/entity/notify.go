package entity

import (
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"time"
)

type Notify struct {
	ID            string      `json:"id" swaggerignore:"true"`
	ChildID       string      `json:"child_id"`
	DestinationID string      `json:"destination_id"`
	FromUserID    string      `json:"from_user_id"`
	FromUser      entity.User `json:"from_user"  gorm:"foreignKey:FromUserID"`
	ToUserID      string      `json:"to_user_id"`
	ToUser        entity.User `json:"to_user" gorm:"foreignKey:ToUserID"`
	Type          string      `json:"type"`
	Title         string      `json:"title"`
	Body          string      `json:"body"`
	IsRead        bool        `json:"is_read" gorm:"default:false"`
	IsImportant   bool        `json:"is_important" gorm:"default:false"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
}
