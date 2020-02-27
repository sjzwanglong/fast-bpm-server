package server

func Register() {
	// 注册流程基本信息路由
	registerProcessDefRouter()

	//注册公司信息路由
	registerOrgCmpyRouter()

	//注册部门信息路由
	registerOrgDeptRouter()

	// 注册用户信息路由
	registerOrgUserRouter()

	// 注册部门用户关系路由
	registerOrgDeptUserRouter()
}
