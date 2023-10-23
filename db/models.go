package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Directory struct {
	gorm.Model

	UUID       string `gorm:"column:UUID;primary_key"`
	Path       string `gorm:"column:path"`
	FolderName string `gorm:"column:folder_name"`
	FileCount  int8   `gorm:"column:file_count"`
}

type Movies struct {
	gorm.Model

	RefId    string `gorm:"column:ref_id"`
	ShowName string `gorm:"column:show_name"`
}

type Shows struct {
	gorm.Model

	RefId    string `gorm:"column:ref_id"`
	ShowName string `gorm:"column:show_name"`
}
