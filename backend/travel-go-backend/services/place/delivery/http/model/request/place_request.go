package request

type CreatePlaceReq struct {
	Thumbnail   string                     `json:"thumbnail"`
	Name        string                     `json:"name"`
	PlaceTypeID string                     `json:"place_type_id"`
	Lat         float64                    `json:"lat"`
	Lng         float64                    `json:"lng"`
	Address     string                     `json:"address"`
	Img         []CreatePlaceImgReq        `json:"img"`
	ImgCheckin  []CreatePlaceImgCheckinReq `json:"img_checkin"`
}

type CreatePlaceImgReq struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type CreatePlaceImgCheckinReq struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type UpdatePlaceReq struct {
	Thumbnail   string                     `json:"thumbnail"`
	Name        string                     `json:"name"`
	PlaceTypeID string                     `json:"place_type_id"`
	Lat         float64                    `json:"lat"`
	Lng         float64                    `json:"lng"`
	Address     string                     `json:"address"`
	Img         []UpdatePlaceImgReq        `json:"img"`
	ImgCheckin  []UpdatePlaceImgCheckinReq `json:"img_checkin"`
}

type UpdatePlaceImgReq struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type UpdatePlaceImgCheckinReq struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type ListReviewPlaceReq struct {
	Limit   int    `form:"limit"`
	Page    int    `form:"page"`
	PlaceID string `json:"place_id"`
}

type ListChildReviewReq struct {
	Limit    int    `form:"limit"`
	Page     int    `form:"page"`
	ReviewID string `json:"review_id"`
}

type ListPlaceReq struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

type FavoritePlaceAddReq struct {
	IsFavorite bool   `json:"is_favorite"`
	PlaceID    string `json:"place_id"`
	UserID     string `json:"user_id" swaggerignore:"true"`
}

type ListFavoritePlaceReq struct {
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
	UserID string `json:"user_id" swaggerignore:"true"`
}
