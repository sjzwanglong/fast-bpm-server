package server

import (
	"fast-bpm/controller"
)

func init() {
	Register(registerOrgDeptRouter)
}

func registerOrgDeptRouter() {
	ctl := &controller.OrgDeptController{} //controller对象
	dualGroup := r.Group("depts")
	singularGroup := r.Group("dept")
	{
		addDualBaseRouter(dualGroup, ctl)
		addSingularBaseRouter(singularGroup, ctl)
	}
}
