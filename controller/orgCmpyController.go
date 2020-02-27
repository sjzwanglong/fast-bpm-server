package controller

import "fast-bpm/service"

type OrgCmpyController struct {
	BaseController
}

var orgCmpyServ *service.OrgCmpyServ = new(service.OrgCmpyServ)
