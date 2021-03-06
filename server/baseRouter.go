package server

import (
	"fast-bpm/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加集合对象路由组
func addDualBaseRouter(instanceFunc func() interface{}, group *gin.RouterGroup, ctl controller.Controller) {
	group.GET("/", func(c *gin.Context) {
		list := instanceFunc()
		ctl.List(list, c.Query("query"), c.Query("order"))
		c.JSON(200, list)
	})
	group.GET("/page/:pageCount/:pageNo", func(c *gin.Context) {
		list := instanceFunc()
		pageCount, _ := strconv.Atoi(c.Param("pageCount"))
		pageNo, _ := strconv.Atoi(c.Param("pageNo"))
		ctl.Page(list, pageCount, pageNo, c.Query("query"), c.Query("order"))
		c.JSON(200, list)
	})
}

// 添加单体对象路由组
func addSingularBaseRouter(instanceFunc func() interface{}, group *gin.RouterGroup, ctl controller.Controller) {
	group.GET("/:id", func(c *gin.Context) {
		out := instanceFunc()
		ctl.ById(out, c.Param("id"))
		c.JSON(200, out)
	})
	group.POST("/", func(c *gin.Context) {
		out := instanceFunc()
		err := c.ShouldBindJSON(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
			return
		}
		ctl.Create(out)
		c.JSON(200, out)
	})
	group.PUT("/:id", func(c *gin.Context) {
		out := instanceFunc()
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("URL绑定参数失败！ err=", err)
			return
		}
		err = c.ShouldBindJSON(out)
		if err != nil {
			fmt.Println("JSON绑定参数失败！ err=", err)
			return
		}
		ctl.Update(out)
		c.JSON(200, out)
	})
	group.DELETE("/delete/:id", func(c *gin.Context) {
		out := instanceFunc()
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
			return
		}
		ctl.Delete(out)
		c.JSON(200, out)
	})
	group.DELETE("/destroy/:id", func(c *gin.Context) {
		out := instanceFunc()
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
			return
		}
		c.Param("destroy")
		ctl.Destroy(out)
		c.JSON(200, out)
	})
}
