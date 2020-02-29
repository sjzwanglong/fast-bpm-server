package controller

import (
	"fast-bpm/model"
	"fast-bpm/service"
)

type LoginController struct {
	BaseController
}

func (l *LoginController) Login(loginName, password string) *model.OrgUserModel {
	var orgUserServ *service.OrgUserServ = new(service.OrgUserServ)
	return orgUserServ.Login(loginName, password)
}
