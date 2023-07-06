package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/delivery/http/model/request"
	place_en "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/place/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/entity"
	reviewus "gitlab.com/virtual-travel/travel-go-backend/services/place/domain/review/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_request"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"net/http"
)

type PlaceHandler struct {
	PlaceUseCase  usecase.PlaceUseCase
	ReviewUseCase reviewus.ReviewUseCase
}

func NewPlaceHandler(accountUseCase usecase.PlaceUseCase, review reviewus.ReviewUseCase) PlaceHandler {
	return PlaceHandler{
		PlaceUseCase:  accountUseCase,
		ReviewUseCase: review,
	}
}

// @Summary Api tạo địa điểm
// @Tags place
// @description Api tạo địa điểm
// @Accept  json
// @Produce  json
// @Param data body request.CreatePlaceReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/create [post]
func (p PlaceHandler) CreatePlace(c *gin.Context) {
	body := request.CreatePlaceReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := p.PlaceUseCase.CreatePlace(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo địa điểm thành công"))
}

// @Summary Api cập nhật địa điểm
// @Tags place
// @description Api cập nhật địa điểm
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param data body request.UpdatePlaceReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/update/{id} [patch]
func (p PlaceHandler) UpdatePlace(c *gin.Context) {
	id := c.Param("id")
	body := request.UpdatePlaceReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	placeImg := []place_en.PlaceImg{}
	for i := range body.Img {
		imgItem := place_en.PlaceImg{
			ID:      randutil.UUIDRand(),
			PlaceID: id,
			Name:    body.Img[i].Name,
			Url:     body.Img[i].Url,
		}
		placeImg = append(placeImg, imgItem)
	}

	placeImgCheckin := []place_en.PlaceImgCheckin{}
	for i := range body.Img {
		imgItemCheckin := place_en.PlaceImgCheckin{
			ID:      randutil.UUIDRand(),
			PlaceID: id,
			Name:    body.ImgCheckin[i].Name,
			Url:     body.ImgCheckin[i].Url,
		}
		placeImgCheckin = append(placeImgCheckin, imgItemCheckin)
	}

	place := place_en.Place{
		ID:              id,
		Thumbnail:       body.Thumbnail,
		Name:            body.Name,
		PlaceTypeID:     &body.PlaceTypeID,
		Lat:             body.Lat,
		Lng:             body.Lng,
		Address:         body.Address,
		PlaceImg:        placeImg,
		PlaceImgCheckin: placeImgCheckin,
	}
	if body.PlaceTypeID == "" {
		place.PlaceTypeID = nil
	}

	err := p.PlaceUseCase.UpdatePlace(c, place)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Cập nhật địa điểm thành công"))
}

// @Summary Api lấy địa điểm
// @Tags place
// @description Api lấy địa điểm
// @Accept  json
// @Produce  json
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/list [get]
func (p PlaceHandler) GetPlace(c *gin.Context) {
	places, err := p.PlaceUseCase.GetPlace(c)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}

// @Summary Api lấy địa điểm cho cms
// @Tags cms
// @description Api lấy địa điểm cho cms
// @Accept  json
// @Produce  json
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /cms/place/list [get]
func (p PlaceHandler) GetPlacePaging(c *gin.Context) {
	body := request.ListPlaceReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	places, err := p.PlaceUseCase.GetPlacePaging(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}

// @Summary Api lấy chi tiết địa điểm
// @Tags place
// @description Api lấy chi tiết địa điểm
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/detail/{id} [get]
func (p PlaceHandler) GetDetailPlace(c *gin.Context) {
	id := c.Param("id")
	place, err := p.PlaceUseCase.GetDetailPlace(c, id)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(place, ""))
}

// @Summary Api tạo đánh giá địa điểm
// @Tags review
// @description Api tạo đánh giá địa điểm
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param data body entity.Review true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /review/create [post]
func (p PlaceHandler) CreateReviewPlace(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := entity.Review{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.UserID = &userId

	err := p.ReviewUseCase.CreateReview(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo địa điểm thành công"))
}

// @Summary Api tạo phản hồi đánh giá
// @Tags review
// @description Api tạo phản hồi đánh giá
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param parent_id path string true "parent_id"
// @Param data body entity.ChildReview true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /review/child/create/{parent_id} [post]
func (p PlaceHandler) CreateChildReviewPlace(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)
	parentID := c.Param("parent_id")

	body := entity.ChildReview{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.UserID = &userId
	body.ParentID = parentID

	err := p.ReviewUseCase.CreateChildReview(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo phản hồi đánh giá"))
}

// @Summary Api lấy đánh giá địa điểm
// @Tags review
// @description Api lấy đánh giá địa điểm
// @Accept  json
// @Produce  json
// @Param place_id path string true "place_id"
// @Param limit query string false "Limit record for getting"
// @Param page query string false "Current page"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /review/list/{place_id} [get]
func (p PlaceHandler) GetReviewPlace(c *gin.Context) {
	placeID := c.Param("place_id")
	body := request.ListReviewPlaceReq{}
	if err := c.Bind(&body); err != nil {
		message := api_request.ValidateReq(c, err)
		c.JSON(http.StatusOK, api_response.SimpleErrorResponse(http.StatusBadRequest, message...))
		return
	}
	body.PlaceID = placeID
	places, err := p.ReviewUseCase.GetReviewByPlaceID(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}

// @Summary Api lấy danh sách đánh giá con
// @Tags review
// @description Api lấy danh sách đánh giá con
// @Accept  json
// @Produce  json
// @Param review_id path string true "review_id"
// @Param limit query string false "Limit record for getting"
// @Param page query string false "Current page"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /child-review/list/{review_id} [get]
func (p PlaceHandler) GetChildReview(c *gin.Context) {
	reviewID := c.Param("review_id")
	body := request.ListChildReviewReq{}
	if err := c.Bind(&body); err != nil {
		message := api_request.ValidateReq(c, err)
		c.JSON(http.StatusOK, api_response.SimpleErrorResponse(http.StatusBadRequest, message...))
		return
	}
	body.ReviewID = reviewID
	places, err := p.ReviewUseCase.GetChildReviewByReviewID(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}

// @Summary Api thêm yêu thích địa điểm
// @Tags favorite
// @description Api thêm yêu thích địa điểm
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param data body request.FavoritePlaceAddReq true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/favorite [post]
func (p PlaceHandler) AddFavoritePlace(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := request.FavoritePlaceAddReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.UserID = userId

	err := p.PlaceUseCase.FavoritePlace(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	mess := "Yêu thích địa điểm thành công"
	if body.IsFavorite == false {
		mess = "Bỏ yêu thích địa điểm thành công"
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, mess))
}

// @Summary Api lấy địa điểm yêu thích
// @Tags favorite
// @description Api lấy địa điểm yêu thích
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param limit query string false "Limit record for getting"
// @Param page query string false "Current page"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /place/favorite/list [get]
func (p PlaceHandler) GetFavoritePlace(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := request.ListFavoritePlaceReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.UserID = userId

	places, err := p.PlaceUseCase.GetFavoritePlacePaging(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(places, ""))
}
