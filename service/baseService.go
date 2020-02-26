package service

import "fast-bpm/dao"

type BaseServ struct {
}

var baseDao *dao.BaseDao = new(dao.BaseDao)

func (bs *BaseServ) ById(out interface{}, id string) {
	baseDao.ById(out, id)
}

func (bs *BaseServ) List(out interface{}, where string, params ...interface{}) {
	if len(params) > 0 {
		baseDao.List(out, where, params)
	} else {
		baseDao.List(out, where)
	}
}

func (bs *BaseServ) Page(out interface{}, limit int, offset int, where string, params ...interface{}) {
	if len(params) > 0 {
		baseDao.Page(out, limit, offset, where, params)
	} else {
		baseDao.Page(out, limit, offset, where)
	}
}

func (bs *BaseServ) Create(out interface{}) {
	baseDao.Create(out)
}

func (bs *BaseServ) Update(out interface{}) {
	baseDao.Update(out)
}

func (bs *BaseServ) Delete(out interface{}) {
	baseDao.Delete(out)
}

func (bs *BaseServ) Destroy(out interface{}) {
	baseDao.Destroy(out)
}
