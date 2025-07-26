package config

import (
	"errors"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger //Logger pointer
)

func Init() error {
	return errors.New("Init Failed")
}

func GetLogger(prefix string) *Logger {
	//Initialize logger
	logger := NewLogger(prefix)
	return logger
}
