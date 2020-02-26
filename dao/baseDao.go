package dao

import (
	"fast-bpm/db"
)

type BaseDao struct {
}

func (bd *BaseDao) ById(out interface{}, id string) {
	db.GetDBFactory().Pool.First(out, "ID = ?", id)
}

func (bd *BaseDao) List(out interface{}, where string, params ...interface{}) {
	if len(params) > 0 {
		db.GetDBFactory().Pool.Find(out, where, params)
	} else {
		db.GetDBFactory().Pool.Find(out, where)
	}
}

func (bd *BaseDao) Page(out interface{}, limit int, offset int, where string, params ...interface{}) {
	if len(params) > 0 {
		db.GetDBFactory().Pool.Limit(limit).Offset(offset).Find(out, where, params)
	} else {
		db.GetDBFactory().Pool.Limit(limit).Offset(offset).Find(out, where)
	}
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
