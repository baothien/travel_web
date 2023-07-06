package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
)

var ViDict = []i18n.Message{
	{
		ID:    constant.ERROR_INTERNAL_SERVER,
		Other: "Đã có lỗi xảy ra",
	},
	{
		ID:    constant.BAD_REQ,
		Other: "Yêu cầu không hợp lệ",
	},
	{
		ID:    constant.DATABASE_ERROR,
		Other: "Đã xảy ra sự cố với Cơ sở dữ liệu",
	},
	{
		ID:    constant.RECORD_NOT_FOUND,
		Other: "Bản ghi không tồn tại",
	},
}
