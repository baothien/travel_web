package entity

import (
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"time"
)

type Review struct {
	ID          string        `json:"id" swaggerignore:"true"`
	PlaceID     *string       `json:"place_id"`
	UserID      *string       `json:"user_id" swaggerignore:"true"`
	User        entity.User   `json:"user,omitempty" gorm:"foreignKey:UserID" swaggerignore:"true"`
	Rate        int           `json:"rate"`
	Description string        `json:"description"`
	Like        int           `json:"like" gorm:"-"`
	Dislike     int           `json:"dislike" gorm:"-"`
	Reaction    []entity.User `json:"-" gorm:"many2many:reaction_reviews"`

	ReviewImg   []ReviewImg    `json:"review_img,omitempty" gorm:"foreignKey:ReviewID"`
	ChildReview *[]ChildReview `json:"child_review,omitempty" gorm:"foreignKey:ParentID" swaggerignore:"true"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
}

type ReactionReview struct {
	ReviewID   string `json:"review_id"`
	UserID     string `json:"user_id"`
	ActionType string `json:"action_type"`
}

type ReviewImg struct {
	ID       string `json:"id" swaggerignore:"true"`
	ReviewID string `json:"review_id" swaggerignore:"true"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type ChildReview struct {
	ID          string           `json:"id" swaggerignore:"true"`
	ParentID    string           `json:"parent_id"`
	UserID      *string          `json:"user_id" swaggerignore:"true"`
	User        entity.User      `json:"user,omitempty" gorm:"foreignKey:UserID" swaggerignore:"true"`
	Description string           `json:"description"`
	ReviewImg   []ChildReviewImg `json:"review_img,omitempty" gorm:"foreignKey:ReviewID"`
	CreatedAt   time.Time        `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
}

type ChildReviewImg struct {
	ID       string `json:"id" swaggerignore:"true"`
	ReviewID string `json:"child_review_id" swaggerignore:"true"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}
