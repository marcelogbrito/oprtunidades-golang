package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	//return errors.New("fake error")
	return nil
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger := NewLogger(p)
	return logger
}
