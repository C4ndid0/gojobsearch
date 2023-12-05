package handler

import (
	"github.com/C4ndid0/gojobsearh/config"
	"gorm.io/gorm"
)

var (
  logger *config.Logger
  db *gorm.DB
)

func InitializeHandler(){
  logger = config.GetLogger("handler")
  db = config.GetSQLite() 
}
