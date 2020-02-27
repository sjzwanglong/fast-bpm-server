package controller

import "fast-bpm/service"

type OrgDeptController struct {
	BaseController
}

var orgDeptServ *service.OrgDeptServ = new(service.OrgDeptServ)
