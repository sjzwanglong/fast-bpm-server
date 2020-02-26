package controller

import "fast-bpm/service"

type BaseController struct {
}

type Controller interface {
	ById(out interface{}, id string)
	List(out interface{}, where string, params ...interface{})
	Page(out interface{}, limit int, offset int, where string, params ...interface{})
	Create(out interface{})
	Update(out interface{})
	Delete(out interface{})
	Destroy(out interface{})
}

var baseServ *service.BaseServ = new(service.BaseServ)

func (bc *BaseController) ById(out interface{}, id string) {
	baseServ.ById(out, id)
}

func (bc *BaseController) List(out interface{}, where string, params ...interface{}) {
	if len(params) > 0 {
		baseServ.List(out, where, params)
	} else {
		baseServ.List(out, where)
	}
}

func (bc *BaseController) Page(out interface{}, limit int, offset int, where string, params ...interface{}) {
	if len(params) > 0 {
		baseServ.Page(out, limit, offset, where, params)
	} else {
		baseServ.Page(out, limit, offset, where)
	}

}

func (bc *BaseController) Create(out interface{}) {
	baseServ.Create(out)
}

func (bc *BaseController) Update(out interface{}) {
	baseServ.Update(out)
}

func (bc *BaseController) Delete(out interface{}) {
	baseServ.Delete(out)
}

func (bc *BaseController) Destroy(out interface{}) {
	baseServ.Destroy(out)
}
