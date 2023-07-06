package request

type ListCheckinReq struct {
	Limit   int    `form:"limit"`
	Page    int    `form:"page"`
	PlaceID string `form:"place_id"`
	UserID  string `json:"user_id"`
}
