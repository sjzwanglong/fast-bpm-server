package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func StartServer() {
	err := r.Run()
	if err != nil {
		fmt.Println("HTTP 服务启动失败！ err=", err)
		return
	}
}
