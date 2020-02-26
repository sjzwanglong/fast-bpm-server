package service

import (
	"fast-bpm/dao"
)

type ProcessDefServ struct {
	BaseServ
}

var processDefDao *dao.ProcessDefDao = new(dao.ProcessDefDao)