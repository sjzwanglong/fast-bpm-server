package model

import (
	"fast-bpm/utils"
)

type ProcessDefModel struct {
	BaseModel
	PCode   string `gorm:"column:P_CODE"`
	PName   string `gorm:"column:P_NAME"`
	Version uint   `gorm:"column:VERSION"`
}

func (p ProcessDefModel) TableName() string {
	return utils.ProcessDefInfo
}

func (c ProcessDefModel) Clone() *ProcessDefModel {
	return new(ProcessDefModel)
}

func (c ProcessDefModel) CloneList() *[]*ProcessDefModel {
	return new([]*ProcessDefModel)
}
