package model

import "fast-bpm/utils"

type OrgDeptModel struct {
	BaseModel
	DeptCode        string `gorm:"column:deptCode" form:"deptCode" uri:"deptCode" json:"deptCode"`
	DeptName        string `gorm:"column:deptName" form:"deptName" uri:"deptName" json:"deptName"`
	DeptDescription string `gorm:"column:deptDescription" form:"deptDescription" uri:"deptDescription" json:"deptDescription"`
	DeptType        int    `gorm:"column:deptType" form:"deptType" uri:"deptType" json:"deptType"`
	DeptSort        int    `gorm:"column:deptSort" form:"deptSort" uri:"deptSort" json:"deptSort"`
	DeptLevel       int    `gorm:"column:deptLevel" form:"deptLevel" uri:"deptLevel" json:"deptLevel"`
	DeptParent      string `gorm:"column:deptParent" form:"deptParent" uri:"deptParent" json:"deptParent"`
	CmpyId          string `gorm:"column:cmpyId" form:"cmpyId" uri:"cmpyId" json:"cmpyId"`
}

func (d OrgDeptModel) TableName() string {
	return utils.OrgDeptInfo
}

func (d OrgDeptModel) Clone() interface{} {
	return new(OrgDeptModel)
}

func (d OrgDeptModel) CloneList() interface{} {
	return new([]*OrgDeptModel)
}
