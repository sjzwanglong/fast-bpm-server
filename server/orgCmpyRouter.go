package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func init() {
	Register(registerOrgCmpyRouter)
}

func registerOrgCmpyRouter() {
	ctl := &controller.OrgCmpyController{} //controller对象
	dualGroup := r.Group("cmpys")
	singularGroup := r.Group("cmpy")
	instance := &model.OrgCmpyModel{}
	{
		addDualBaseRouter(instance.CloneList, dualGroup, ctl)
		addSingularBaseRouter(instance.Clone, singularGroup, ctl)
	}
}
