package service

import "fast-bpm/dao"

type OrgCmpyServ struct {
	BaseServ
}

var orgCmpyDao *dao.OrgCmpyDao = new(dao.OrgCmpyDao)
