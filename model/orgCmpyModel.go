package model

import "fast-bpm/utils"

type OrgCmpyModel struct {
	BaseModel
	CmpyCode        string `gorm:"column:cmpyCode" form:"cmpyCode" uri:"cmpyCode" json:"cmpyCode"`
	CmpyName        string `gorm:"column:cmpyName" form:"cmpyName" uri:"cmpyName" json:"cmpyName"`
	CmpyFullName    string `gorm:"column:cmpyFullName" form:"cmpyFullName" uri:"cmpyFullName" json:"cmpyFullName"`
	CmpyDescription string `gorm:"column:cmpyDescription" form:"cmpyDescription" uri:"cmpyDescription" json:"cmpyDescription"`
	CmpyLogo        string `gorm:"column:cmpyLogo" form:"cmpyLogo" uri:"cmpyLogo" json:"cmpyLogo"`
	CmpyLevel       int    `gorm:"column:cmpyLevel" form:"cmpyLevel" uri:"cmpyLevel" json:"cmpyLevel"`
	CmpySort        int    `gorm:"column:cmpySort" form:"cmpySort" uri:"cmpySort" json:"cmpySort"`
	CmpyParent      string `gorm:"column:cmpyParent" form:"cmpyParent" uri:"cmpyParent" json:"cmpyParent"`
}

func (p OrgCmpyModel) TableName() string {
	return utils.OrgCmpyInfo
}
