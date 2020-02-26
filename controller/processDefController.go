package controller

import "fast-bpm/service"

type ProcessDefController struct {
	BaseController
}

var processDefServ *service.ProcessDefServ = new(service.ProcessDefServ)