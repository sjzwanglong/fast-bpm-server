package server

import (
	"fast-bpm/controller"
)

func registerOrgCmpyRouter() {
	ctl := &controller.OrgCmpyController{} //controller对象
	dualGroup := r.Group("cmpys")
	singularGroup := r.Group("cmpy")
	{
		addDualBaseRouter(dualGroup, ctl)
		addSingularBaseRouter(singularGroup, ctl)
	}
}
