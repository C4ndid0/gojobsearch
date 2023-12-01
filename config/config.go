package config

import (
	"errors"

	"gorm.io/gorm"
)

var(
  db *gorm.DB
  logger *Logger
)

func Init() error {
  return errors.New("fake Error")
}

func GetLogger(p string) *Logger {
  // Initialize logger 
  logger = NewLogger(p)
  return logger 
}
