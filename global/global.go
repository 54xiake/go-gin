package global

import "gorm.io/gorm"

var (
	DB        *gorm.DB
	UploadDir string //上传文件目录
)
