package server

import (
	"fast-bpm/controller"
)

func init() {
	Register(registerProcessDefRouter)
}

func registerProcessDefRouter() {
	ctl := &controller.ProcessDefController{} //controller对象
	dualGroup := r.Group("processDefs")
	singularGroup := r.Group("processDef")
	{
		addDualBaseRouter(dualGroup, ctl)
		addSingularBaseRouter(singularGroup, ctl)
	}
}
