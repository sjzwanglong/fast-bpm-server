package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
)

func registerProcessDefRouter() {
	ctl := &controller.ProcessDefController{}     //controller对象
	processDefModel := &model.ProcessDefModel{}   //model对象
	processDefList := &[]*model.ProcessDefModel{} //list对象
	processDefGroup := r.Group("processDef")
	{
		addBaseRouter(processDefGroup, ctl, processDefModel, processDefList)
	}
	
}
