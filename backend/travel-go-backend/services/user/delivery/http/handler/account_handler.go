package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/user/domain/account/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_request"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"net/http"
)

type AccountHandler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewAccountHandler(accountUseCase usecase.AccountUseCase) AccountHandler {
	return AccountHandler{AccountUseCase: accountUseCase}
}

// @Summary Request register account customer
// @Tags account
// @description Api yêu cầu tạo tài khoản khách hàng
// @Accept  json
// @Produce  json
// @Param data body entity.User true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/register [post]
func (a AccountHandler) RegisterUser(c *gin.Context) {
	body := entity.User{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := a.AccountUseCase.CreateAccount(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, "Tạo tài khoản thành công"))
}

// @Summary Api đăng nhập
// @Tags account
// @description Api đăng nhập
// @Accept  json
// @Produce  json
// @Param data body request.LoginReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/login [post]
func (a AccountHandler) Login(c *gin.Context) {
	body := request.LoginReq{}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := a.AccountUseCase.Login(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, ""))
}

// @Summary Api get profile
// @Tags account
// @description Api get profile
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/profile [get]
func (a AccountHandler) Profile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	user, err := a.AccountUseCase.Profile(c, userId)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, ""))
}

// @Summary API cập nhật thông tin tài khoản
// @Tags account
// @description API cập nhật thông tin tài khoản
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param data body request.UpdateUserReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/profile [patch]
func (a AccountHandler) UpdateProfile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := request.UpdateUserReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user := entity.User{
		ID:          userId,
		FullName:    body.FullName,
		UserName:    body.Email,
		Email:       body.Email,
		Phone:       body.Phone,
		DateOfBirth: body.DateOfBirth,
		Avatar:      body.Avatar,
	}

	user, err := a.AccountUseCase.UpdateAccount(c, user)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, "Cập nhật thành công"))
}

// @Summary api refresh token for customer
// @Tags account
// @description Api refresh token cho người dùng
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/refresh-token [get]
func (s AccountHandler) RefreshToken(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	userReq := entity.User{
		ID: userId,
	}

	userInfo, err := s.AccountUseCase.RefreshToken(c, userReq)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(userInfo, ""))
}

// @Summary Api chi tiết người dùng
// @Tags account
// @description Api chi tiết người dùng
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "id"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/detail/{id} [get]
func (a AccountHandler) DetailUser(c *gin.Context) {
	userId := c.Param("id")

	user, err := a.AccountUseCase.Profile(c, userId)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, ""))
}

// @Summary Api lấy danh sách người dùng
// @Tags account
// @description Api lấy danh sách người dùng
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param limit query string false "Limit record for getting"
// @Param page query string false "Current page"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/list [get]
func (s AccountHandler) ListAccount(c *gin.Context) {
	body := request.ListUser{}
	if err := c.Bind(&body); err != nil {
		message := api_request.ValidateReq(c, err)
		c.JSON(http.StatusOK, api_response.SimpleErrorResponse(http.StatusBadRequest, message...))
		return
	}
	//
	user, err := s.AccountUseCase.ListUser(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusOK, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(user, ""))
}

// @Summary API đổi mật khẩu
// @Tags account
// @description API đổi mật khẩu
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param data body request.ChangePassReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /user/change-pass [patch]
func (a AccountHandler) ChangePassword(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := request.ChangePassReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.ID = userId

	err := a.AccountUseCase.ChangePass(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(nil, "Cập nhật mật khẩu thành công"))
}
