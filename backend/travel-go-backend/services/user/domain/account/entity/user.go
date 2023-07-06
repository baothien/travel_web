package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type User struct {
	ID          string    `json:"id" swaggerignore:"true"`
	UserName    string    `json:"user_name" gorm:"unique" swaggerignore:"true"`
	UserType    string    `json:"user_type" swaggerignore:"true"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	DateOfBirth string    `json:"date_of_birth"`
	Avatar      string    `json:"avatar"`
	Token       *Token    `json:"token,omitempty" gorm:"-" swaggerignore:"true"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
}

type Token struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiredTime  time.Duration `json:"expired_time"`
}

type JwtCustomClaims struct {
	UserId   string
	UserType string
	DeviceId string
	jwt.StandardClaims
}
