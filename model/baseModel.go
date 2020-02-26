package model

import (
	"fast-bpm/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type CommonModel interface {
	TableName()
}

type BaseModel struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:CREATED_AT"`
	UpdatedAt time.Time  `gorm:"column:UPDATED_AT"`
	DeletedAt *time.Time `gorm:"column:DELETED_AT",sql:"index"`
}

func (bd *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", utils.GetUUID())
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (bd *BaseModel) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
