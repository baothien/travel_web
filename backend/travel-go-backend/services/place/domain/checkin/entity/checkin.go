package entity

import (
	placeen "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"time"
)

type Checkin struct {
	ID        string         `json:"id" swaggerignore:"true"`
	PlaceID   *string        `json:"place_id"`
	Place     *placeen.Place `json:"place" gorm:"foreignKey:PlaceID"`
	UserID    *string        `json:"user_id" swaggerignore:"true"`
	User      entity.User    `json:"user,omitempty" gorm:"foreignKey:UserID" swaggerignore:"true"`
	Name      string         `json:"name"`
	Url       string         `json:"url"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
}
