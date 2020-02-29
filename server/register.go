package server

// 调用传入的路由注册方法
func Register(reg func()) {
	reg()
}
