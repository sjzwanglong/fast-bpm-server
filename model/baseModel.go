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
	ID        string     `gorm:"primary_key" form:"id" uri:"id" json:"id"`
	CreatedBy string     `gorm:"column:createdBy"`
	CreatedAt time.Time  `gorm:"column:createdAt"`
	UpdatedAt time.Time  `gorm:"column:createdAt"`
	DeletedAt *time.Time `gorm:"column:deletedAt",sql:"index"`
	SysState  int        `gorm:"column:sysState" form:"sysState" uri:"sysState" json:"sysState"`
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
