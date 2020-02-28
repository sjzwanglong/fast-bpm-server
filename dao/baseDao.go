package dao

import (
	"fast-bpm/db"
)

type BaseDao struct {
}

func (bd *BaseDao) ById(out interface{}, id string) {
	db.GetDBFactory().First(out, "ID = ?", id)
}

func (bd *BaseDao) List(out interface{}, params interface{}, order string) {
	if len(order) <= 0 {
		order = db.GetDBConfig().DefaultOrder
	}
	if params != nil {
		db.GetDBFactory().Where(params).Order(order).Find(out)
	} else {
		db.GetDBFactory().Order(order).Find(out)
	}
}

func (bd *BaseDao) Page(out interface{}, limit, offset int, params interface{}, order string) {
	if len(order) <= 0 {
		order = db.GetDBConfig().DefaultOrder
	}
	if params != nil {
		db.GetDBFactory().Where(params).Order(order).Limit(limit).Offset(offset).Find(out)
	} else {
		db.GetDBFactory().Order(order).Limit(limit).Offset(offset).Find(out)
	}
}

func (bd *BaseDao) Create(out interface{}) {
	db.GetDBFactory().Create(out)
}

func (bd *BaseDao) Update(out interface{}) {
	db.GetDBFactory().Model(out).Updates(out)
}

func (bd *BaseDao) Delete(out interface{}) {
	db.GetDBFactory().Delete(out)
}

func (bd *BaseDao) Destroy(out interface{}) {
	db.GetDBFactory().Unscoped().Delete(out)
}
