package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func init() {
	Register(registerProcessDefRouter)
}

func registerProcessDefRouter() {
	ctl := &controller.ProcessDefController{} //controller对象
	dualGroup := r.Group("processDefs")
	singularGroup := r.Group("processDef")
	instance := &model.ProcessDefModel{}
	{
		addDualBaseRouter(instance.CloneList, dualGroup, ctl)
		addSingularBaseRouter(instance.Clone, singularGroup, ctl)
	}
}
