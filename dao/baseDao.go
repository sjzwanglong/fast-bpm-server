package dao

import (
	"fast-bpm/db"
)

type BaseDao struct {
}

func (bd *BaseDao) ById(out interface{}, id string) {
	db.GetDBFactory().Pool.First(out, "ID = ?", id)
}

func (bd *BaseDao) List(out interface{}, params interface{}) {
	db.GetDBFactory().Pool.Where(params).Find(out)
}

func (bd *BaseDao) Page(out interface{}, limit, offset int, params interface{}) {
	db.GetDBFactory().Pool.Where(params).Limit(limit).Offset(offset).Find(out)
}

func (bd *BaseDao) Create(out interface{}) {
	db.GetDBFactory().Pool.Create(out)
}

func (bd *BaseDao) Update(out interface{}) {
	db.GetDBFactory().Pool.Model(out).Updates(out)
}

func (bd *BaseDao) Delete(out interface{}) {
	db.GetDBFactory().Pool.Delete(out)
}

func (bd *BaseDao) Destroy(out interface{}) {
	db.GetDBFactory().Pool.Unscoped().Delete(out)
}
