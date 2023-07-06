package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/upload/domain/places/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_request"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"gitlab.com/virtual-travel/travel-go-backend/utils/validatorutil"

	"net/http"
)

type FileHandler struct {
	FileUseCase usecase.FileUseCase
}

func NewFileHandler(accountUseCase usecase.FileUseCase) FileHandler {
	return FileHandler{FileUseCase: accountUseCase}
}

// @Summary Upload file
// @Tags file
// @description Api upload file
// @Accept  json
// @Produce  json
// @Param file formData file true "upload images"
// @Param type formData string true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /file/upload [post]
func (s FileHandler) UploadFile(c *gin.Context) {
	lang := api_request.GetRequestLanguage(c)

	localizer := i18nutil.NewLocalizer(lang)
	body := request.FileReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusOK, api_response.Response{
			ErrorInfo: &api_response.ErrorInfo{
				Code:    constant.BAD_REQ,
				Message: validatorutil.ArrError(err.(validator.ValidationErrors), lang),
			},
			Message: "",
			Data:    nil,
		})
		return
	}

	fileReq, err := c.FormFile("file")
	if err != nil {
		message, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: constant.BAD_REQ,
		})
		c.JSON(http.StatusOK, api_response.Response{
			ErrorInfo: &api_response.ErrorInfo{
				Code:    constant.BAD_REQ,
				Message: []string{message},
			},
			Message: "",
			Data:    nil,
		})
		return
	}
	body.File = fileReq

	fileResp, err := s.FileUseCase.UploadFile(c, body)
	if err != nil {
		message, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: constant.ERROR_INTERNAL_SERVER,
		})
		c.JSON(http.StatusOK, api_response.Response{
			ErrorInfo: &api_response.ErrorInfo{
				Code:    constant.ERROR_INTERNAL_SERVER,
				Message: []string{message},
			},
			Message: "",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, api_response.Response{
		ErrorInfo: nil,
		Message:   "",
		Data:      fileResp,
	})
}
