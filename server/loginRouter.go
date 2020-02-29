package server

import (
	"fast-bpm/controller"
	"fast-bpm/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	Register(registerLoginRouter)
}

func registerLoginRouter() {
	r.POST("/login", login)
	r.POST("/super", superLogin)
}

func login(c *gin.Context) {
	ctl := &controller.LoginController{}
	loginName := c.PostForm("loginName")
	pwd := c.PostForm("password")
	userModel := ctl.Login(loginName, pwd)
	c.JSON(200, userModel)
}

func superLogin(c *gin.Context) {
	loginName := c.PostForm("loginName")
	pwd := c.PostForm("password")

	superAdmin := utils.GetSuperAdmin()
	if superAdmin.UserName == loginName && superAdmin.Password == pwd {
		fmt.Println("登录成功！！！")
	}
	c.JSON(200, superAdmin)
}
