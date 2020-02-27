package controller

import "fast-bpm/service"

type OrgUserController struct {
	BaseController
}

var orgUserServ *service.OrgUserServ = new(service.OrgUserServ)
