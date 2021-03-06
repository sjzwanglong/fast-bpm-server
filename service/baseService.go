package service

import "fast-bpm/dao"

type BaseServ struct {
}

var baseDao *dao.BaseDao = new(dao.BaseDao)

func (bs *BaseServ) ById(out interface{}, id string) {
	baseDao.ById(out, id)
}

func (bs *BaseServ) List(out interface{}, params interface{}, order string) {
	baseDao.List(out, params, order)
}

func (bs *BaseServ) Page(out interface{}, limit, offset int, params interface{}, order string) {
	baseDao.Page(out, limit, offset, params, order)
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
