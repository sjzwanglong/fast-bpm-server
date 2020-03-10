package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func init() {
	Register(registerOrgDeptRouter)
}

func registerOrgDeptRouter() {
	ctl := &controller.OrgDeptController{} //controller对象
	dualGroup := r.Group("depts")
	singularGroup := r.Group("dept")
	instance := &model.OrgDeptModel{}
	{
		addDualBaseRouter(instance.CloneList, dualGroup, ctl)
		addSingularBaseRouter(instance.Clone, singularGroup, ctl)
	}
}
