package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func init() {
	Register(registerOrgDeptUserRouter)
}

func registerOrgDeptUserRouter() {
	ctl := &controller.OrgDeptUserController{} //controller对象
	dualGroup := r.Group("deptusers")
	singularGroup := r.Group("deptuser")
	instance := &model.OrgDeptUserModel{}
	{
		addDualBaseRouter(instance.CloneList, dualGroup, ctl)
		addSingularBaseRouter(instance.Clone, singularGroup, ctl)
	}
}
