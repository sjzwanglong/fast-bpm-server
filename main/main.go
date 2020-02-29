package main

import (
	"fast-bpm/db"
	"fast-bpm/server"
	"fast-bpm/utils"
)

func main() {
	// 初始化超级管理员信息
	utils.InitSuperAdmin()
	// 初始化数据库连接
	db.InitFactory("mysql")
	// 启动web服务器
	server.StartServer()
}
