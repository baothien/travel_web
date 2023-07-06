package request

import "mime/multipart"

type FileReq struct {
	File *multipart.FileHeader `swaggerignore:"true"`
	Type string                `json:"type" form:"type"`
}
