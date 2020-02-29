package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 超级管理员信息
type SuperAdmin struct {
	UserName string `ini:"username" json:"userName"`
	Password string `ini:"password" json:"password"`
}

var superAdmin = new(SuperAdmin)

func InitSuperAdmin() {
	cfg, err := ini.Load(GetAbsPath("../config/config.ini"))
	if err != nil {
		fmt.Println("读取配置文件错误！ err=", err)
		return
	}
	cfg.BlockMode = false
	err = cfg.Section("system").MapTo(superAdmin)
	if err != nil {
		fmt.Println("映射结构体错误！ err=", err)
		return
	}
}

func GetSuperAdmin() *SuperAdmin {
	return superAdmin
}
