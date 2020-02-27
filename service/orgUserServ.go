package service

import "fast-bpm/dao"

type OrgUserServ struct {
	BaseServ
}

var orgUserDao *dao.OrgUserDao = new(dao.OrgUserDao)
