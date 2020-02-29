package service

import (
	"fast-bpm/dao"
	"fast-bpm/model"
)

type OrgUserServ struct {
	BaseServ
}

var orgUserDao *dao.OrgUserDao = new(dao.OrgUserDao)

func (u *OrgUserServ) Login(loginName, password string) *model.OrgUserModel {
	return orgUserDao.Login(loginName, password)
}
