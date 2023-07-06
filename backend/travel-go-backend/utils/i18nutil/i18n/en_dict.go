package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/constant"
)

var EnDict = []i18n.Message{
	{
		ID:    constant.ERROR_INTERNAL_SERVER,
		Other: "Server error",
	},
	{
		ID:    constant.BAD_REQ,
		Other: "Bad request",
	},
	{
		ID:    constant.DATABASE_ERROR,
		Other: "Something went wrong with Database",
	},
	{
		ID:    constant.RECORD_NOT_FOUND,
		Other: "Record not found",
	},
}
