package controller

import "fast-bpm/service"

type OrgDeptUserController struct {
	BaseController
}

var orgDeptUserServ *service.OrgDeptUserServ = new(service.OrgDeptUserServ)
