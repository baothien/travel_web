package fileutil

import (
	"github.com/google/uuid"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"go.uber.org/zap"
	"os"
	"strings"
)

func MoveFileUploads(typeFile string, contentType string) (fileName string, savePath string, filePath string) {
	root := "../../public/"
	uploadPath := "files/" + typeFile + "/"
	// Mkdir if not existed
	if _, err := os.Stat(root + uploadPath); os.IsNotExist(err) {
		logger.Warn("Path does not exist", zap.Error(err))
		_ = os.MkdirAll(root+uploadPath, 0755)
	}

	// Rename file
	newUUID, _ := uuid.NewUUID()

	name := newUUID.String()
	// extension := strings.Split(contentType, "/")
	ext := GetExtensionFromContentType(contentType)
	path := uploadPath + name + "." + ext
	return name, root + path, path
}

func GetExtensionFromContentType(contentType string) string {

	extension := strings.Split(contentType, "/")
	var ext = extension[1]
	//if len(extension) < 2 {
	//	if typeFile == "avatar" {
	//		ext = "jpg"
	//	}
	//} else {
	//	ext = extension[1]
	//}
	if ext == "vnd.openxmlformats-officedocument.wordprocessingml.document" {
		ext = "docx"
	}
	if ext == "msword" {
		ext = "doc"
	}
	if ext == "vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		ext = "xlsx"
	}
	if ext == "vnd.ms-excel\n" {
		ext = "xls"
	}
	return ext
}
