package entity

import (
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/entity"
	user_entity "gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"time"
)

type Place struct {
	ID              string            `json:"id"`
	Thumbnail       string            `json:"thumbnail"`
	Name            string            `json:"name"`
	PlaceTypeID     *string           `json:"place_type_id"`
	PlaceType       *PlaceType        `json:"place_type,omitempty" gorm:"foreignKey:PlaceTypeID"`
	Lat             float64           `json:"lat"`
	Lng             float64           `json:"lng"`
	Address         string            `json:"address"`
	PlaceImg        []PlaceImg        `json:"place_img" gorm:"foreignKey:PlaceID"`
	PlaceImgCheckin []PlaceImgCheckin `json:"place_img_checkin" gorm:"foreignKey:PlaceID"`
	UserFavorite    *user_entity.User `json:"-" gorm:"many2many:userplace_favorites;"`
	Review          []entity.Review   `json:"review" gorm:"foreignKey:PlaceID"`
	CreatedAt       time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
}

type PlaceType struct {
	ID   string `json:"id"`
	Code string `json:"code" gorm:"unique;not null"`
	Name string `json:"name"`
}

type PlaceImg struct {
	ID      string `json:"id"`
	PlaceID string `json:"place_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

type PlaceImgCheckin struct {
	ID      string `json:"id"`
	PlaceID string `json:"place_id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
