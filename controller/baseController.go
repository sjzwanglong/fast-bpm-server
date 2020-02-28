package controller

import (
	"encoding/json"
	"fast-bpm/service"
)

type BaseController struct {
}

type Controller interface {
	ById(out interface{}, id string)
	List(out interface{}, params, order string)
	Page(out interface{}, pageCount, pageNo int, params, order string)
	Create(out interface{})
	Update(out interface{})
	Delete(out interface{})
	Destroy(out interface{})
}

var baseServ *service.BaseServ = new(service.BaseServ)

func (bc *BaseController) ById(out interface{}, id string) {
	baseServ.ById(out, id)
}

func (bc *BaseController) List(out interface{}, params, order string) {
	var tempParams interface{}
	json.Unmarshal([]byte(params), &tempParams)
	baseServ.List(out, tempParams, order)
}

func (bc *BaseController) Page(out interface{}, pageCount, pageNo int, params, order string) {
	var tempParams interface{}
	json.Unmarshal([]byte(params), &tempParams)
	baseServ.Page(out, pageCount, (pageNo-1)*pageCount, tempParams, order)

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
