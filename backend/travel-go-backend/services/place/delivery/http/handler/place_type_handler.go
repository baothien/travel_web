package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"net/http"
)

type PlaceTypeHandler struct {
	PlaceTypeUseCase usecase.PlaceTypeUseCase
}

func NewPlaceTypeHandler(accountUseCase usecase.PlaceTypeUseCase) PlaceTypeHandler {
	return PlaceTypeHandler{PlaceTypeUseCase: accountUseCase}
}

// @Summary Api tạo loại địa điểm
// @Tags place-type
// @description Api tạo loại địa điểm
// @Accept  json
// @Produce  json
// @Param data body request.CreatePlaceTypeReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place-type/create [post]
func (p PlaceTypeHandler) CreatePlaceType(c *gin.Context) {
	body := request.CreatePlaceTypeReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := p.PlaceTypeUseCase.CreatePlaceType(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo loại địa điểm thành công"))
}

// @Summary Api lấy loại địa điểm
// @Tags place-type
// @description Api lấy loại địa điểm
// @Accept  json
// @Produce  json
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place-type/list [get]
func (p PlaceTypeHandler) GetPlaceType(c *gin.Context) {
	places, err := p.PlaceTypeUseCase.GetPlaceType(c)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}
