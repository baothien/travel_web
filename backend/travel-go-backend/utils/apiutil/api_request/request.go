package api_request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/virtual-travel/travel-go-backend/utils/validatorutil"
	"strings"
)

func GetRequestLanguage(c *gin.Context) string {
	if l := c.GetHeader("Accept-Language"); l != "" {
		s := strings.ToLower(l)
		if s == "vie" {
			return "vi"
		}
		return s
	}

	return "en"
}

func GetIPAddress(c *gin.Context) string {
	forwarded := c.GetHeader("X-Forwarded-For")
	if forwarded != "" {
		if strings.Contains(forwarded, ",") {
			return strings.TrimSpace(strings.Split(forwarded, ",")[0])
		}
		return forwarded
	}
	forwarded = c.GetHeader("X-Real-Ip")
	if forwarded != "" {
		if strings.Contains(forwarded, ",") {
			return strings.TrimSpace(strings.Split(forwarded, ",")[0])
		}
		return forwarded
	}
	return c.Request.RemoteAddr
}

func GetDeviceId(c *gin.Context) string {
	deviceId := c.GetHeader("X-Client-Device-Id")
	return deviceId
}

/** Mobile pass this header*/
func GetServerVersion(c *gin.Context) string {
	serverVersion := c.GetHeader("X-Server-Version")
	return serverVersion
}

func GetSDKVersion(c *gin.Context) string {
	sdkVersion := c.GetHeader("sdk-version")
	return sdkVersion
}

func ValidateReq(c *gin.Context, err error) []string {
	lang := GetRequestLanguage(c)
	return validatorutil.ArrError(err.(validator.ValidationErrors), lang)
}
