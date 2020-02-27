package server

import (
	"fast-bpm/controller"
)

func registerOrgUserRouter() {
	ctl := &controller.OrgUserController{} //controller对象
	dualGroup := r.Group("users")
	singularGroup := r.Group("user")
	{
		addDualBaseRouter(dualGroup, ctl)
		addSingularBaseRouter(singularGroup, ctl)
	}
}
