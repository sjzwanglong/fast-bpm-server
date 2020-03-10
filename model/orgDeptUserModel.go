package model

import "fast-bpm/utils"

type OrgDeptUserModel struct {
	BaseModel
	UserId   string `gorm:"column:userId" form:"userId" uri:"userId" json:"userId"`
	DeptId   string `gorm:"column:deptId" form:"deptId" uri:"deptId" json:"deptId"`
	UserPost string `gorm:"column:userPost" form:"userPost" uri:"userPost" json:"userPost"`
}

func (du OrgDeptUserModel) TableName() string {
	return utils.OrgDeptUserInfo
}

func (du OrgDeptUserModel) Clone() interface{} {
	return new(OrgDeptUserModel)
}

func (du OrgDeptUserModel) CloneList() interface{} {
	return new([]*OrgDeptUserModel)
}