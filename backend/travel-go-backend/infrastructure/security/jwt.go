package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"strconv"
	"time"
)

func GenToken(user entity.User) (string, error) {
	jwtExpConfig := viper.Get("JWT_EXPIRED").(string)
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)

	jwtExpDuration :=
		time.Hour * time.Duration(jwtExpValue)

	secretKey := viper.Get("JWT_SECRET_KEY").(string)
	fmt.Println(jwtExpConfig)

	claims := &entity.JwtCustomClaims{
		UserId:   user.ID,
		UserType: user.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return result, nil
}

func GenRefreshToken(user entity.User) (string, error) {
	jwtExpConfig := viper.Get("JWT_REFRESH_EXPIRED").(string)
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)

	jwtExpDuration :=
		time.Hour * time.Duration(jwtExpValue)

	secretRefreshKey := viper.Get("JWT_REFRESH_SECRET_KEY").(string)

	claims := &entity.JwtCustomClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secretRefreshKey))
	if err != nil {
		return "", err
	}
	return result, nil
}

func GenTokenObj(user entity.User) (entity.Token, error) {
	//generate token
	var token entity.Token

	jwtExpConfig := viper.Get("JWT_EXPIRED").(string)
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)
	jwtExpDuration := time.Hour * time.Duration(jwtExpValue) / time.Second
	token.ExpiredTime = jwtExpDuration

	var err error
	token.AccessToken, err = GenToken(user)
	if err != nil {
		return token, err
	}
	token.RefreshToken, err = GenRefreshToken(user)
	if err != nil {
		return token, err
	}

	return token, err
}
