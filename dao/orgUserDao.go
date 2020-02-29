package dao

import (
	"fast-bpm/db"
	"fast-bpm/model"
)

type OrgUserDao struct {
	BaseDao
}

func (u *OrgUserDao) Login(loginName, password string) *model.OrgUserModel {
	userModel := &model.OrgUserModel{}
	db.GetDBFactory().Where("userLoginName = ? and userPassword = ?", loginName, password).First(userModel)
	return userModel
}
