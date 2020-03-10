package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func init() {
	Register(registerOrgUserRouter)
}

func registerOrgUserRouter() {
	ctl := &controller.OrgUserController{} //controller对象
	dualGroup := r.Group("users")
	singularGroup := r.Group("user")
	instance := &model.OrgUserModel{}
	{
		addDualBaseRouter(instance.CloneList, dualGroup, ctl)
		addSingularBaseRouter(instance.Clone, singularGroup, ctl)
	}
}
