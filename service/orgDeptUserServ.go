package service

import "fast-bpm/dao"

type OrgDeptUserServ struct {
	BaseServ
}

var orgDeptUserDao *dao.OrgDeptUserDao = new(dao.OrgDeptUserDao)
