package handler

import (
	"github.com/gustasousagh/jobs-api-golang/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}
