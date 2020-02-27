package server

import (
	"fast-bpm/controller"
	"fast-bpm/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加集合对象路由组
func addDualBaseRouter(group *gin.RouterGroup, ctl controller.Controller) {
	group.GET("/", func(c *gin.Context) {
		list := model.NewModelInstance(group.BasePath())
		ctl.List(list, c.Query("query"))
		c.JSON(200, list)
	})
	group.GET("/page/:limit/:offset", func(c *gin.Context) {
		list := model.NewModelInstance(group.BasePath())
		limit, _ := strconv.Atoi(c.Param("limit"))
		offset, _ := strconv.Atoi(c.Param("offset"))
		ctl.Page(list, limit, offset, c.Query("query"))
		c.JSON(200, list)
	})
}

// 添加单体对象路由组
func addSingularBaseRouter(group *gin.RouterGroup, ctl controller.Controller) {
	group.GET("/:id", func(c *gin.Context) {
		out := model.NewModelInstance(group.BasePath())
		ctl.ById(out, c.Param("id"))
		c.JSON(200, out)
	})
	group.POST("/", func(c *gin.Context) {
		out := model.NewModelInstance(group.BasePath())
		err := c.ShouldBindJSON(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
		}
		ctl.Create(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.PUT("/:id", func(c *gin.Context) {
		out := model.NewModelInstance(group.BasePath())
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("URL绑定参数失败！ err=", err)
		}
		err = c.ShouldBindJSON(out)
		if err != nil {
			fmt.Println("JSON绑定参数失败！ err=", err)
		}
		ctl.Update(out)
		c.JSON(200, out)
	})
	group.DELETE("/delete/:id", func(c *gin.Context) {
		out := model.NewModelInstance(group.BasePath())
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
		}
		ctl.Delete(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.DELETE("/destroy/:id", func(c *gin.Context) {
		out := model.NewModelInstance(group.BasePath())
		err := c.ShouldBindUri(out)
		if err != nil {
			fmt.Println("绑定参数失败！ err=", err)
		}
		c.Param("destroy")
		ctl.Delete(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
