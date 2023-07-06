package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/checkin/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_request"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"net/http"
)

type CheckinHandler struct {
	CheckinUseCase usecase.CheckinUseCase
}

func NewCheckinHandler(accountUseCase usecase.CheckinUseCase) CheckinHandler {
	return CheckinHandler{
		CheckinUseCase: accountUseCase,
	}
}

// @Summary Api lưu checkin
// @Tags checkin
// @description Api lưu checkin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param data body entity.Checkin true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /checkin/create [post]
func (ch CheckinHandler) CreateCheckin(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := entity.Checkin{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	body.UserID = &userId

	err := ch.CheckinUseCase.CreateCheckin(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo checkin thành công"))
}

// @Summary Api lấy danh sách checkin
// @Tags checkin
// @description Api lấy danh sách checkin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param place_id query string false "place_id"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /checkin/list [get]
func (ch CheckinHandler) GetCheckin(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)
	body := request.ListCheckinReq{}
	if err := c.Bind(&body); err != nil {
		message := api_request.ValidateReq(c, err)
		c.JSON(http.StatusOK, api_response.SimpleErrorResponse(http.StatusBadRequest, message...))
		return
	}

	body.UserID = userId

	places, err := ch.CheckinUseCase.GetCheckinPaging(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}
