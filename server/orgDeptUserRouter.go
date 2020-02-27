package server

import (
	"fast-bpm/controller"
)

func registerOrgDeptUserRouter() {
	ctl := &controller.OrgDeptUserController{} //controller对象
	dualGroup := r.Group("deptusers")
	singularGroup := r.Group("deptuser")
	{
		addDualBaseRouter(dualGroup, ctl)
		addSingularBaseRouter(singularGroup, ctl)
	}
}
