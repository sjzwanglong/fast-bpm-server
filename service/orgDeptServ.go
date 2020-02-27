package service

import "fast-bpm/dao"

type OrgDeptServ struct {
	BaseServ
}

var orgDeptDao *dao.OrgDeptDao = new(dao.OrgDeptDao)
