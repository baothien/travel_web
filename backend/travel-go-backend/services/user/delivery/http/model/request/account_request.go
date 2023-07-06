package request

type LoginReq struct {
	UserName string `json:"user_name" binding:"required" example:"huy@mgmail.com"`
	Password string `json:"password" binding:"required" example:"lPbtgmKE1aWsBTsb"`
}

type ListUser struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

type UpdateUserReq struct {
	ID          string `json:"id" swaggerignore:"true"`
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Avatar      string `json:"avatar"`
}

type ChangePassReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	ID          string `json:"id" swaggerignore:"true"`
}
