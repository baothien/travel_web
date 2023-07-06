package request

type ListNotifyReq struct {
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
	UserID string `json:"user_id" swaggerignore:"true"`
}

type UpdateIsRead struct {
	ID     string `json:"id"  swaggerignore:"true"`
	IsRead bool   `json:"is_read"`
}
