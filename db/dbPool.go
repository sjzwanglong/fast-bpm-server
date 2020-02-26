package db

import "github.com/jinzhu/gorm"

type DBPool interface {
	initPool(config *Config) *gorm.DB
}
