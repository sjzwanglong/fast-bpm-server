package server

import (
	"fast-bpm/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

func addBaseRouter(group *gin.RouterGroup, ctl controller.Controller, out interface{}, list interface{}) {
	group.GET("/", func(c *gin.Context) {
		ctl.List(list, "")
		c.JSON(200, list)
	})
	group.GET("/page/:limit/:offset", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.Param("limit"))
		offset, _ := strconv.Atoi(c.Param("offset"))
		ctl.Page(list, limit, offset, "")
		c.JSON(200, list)
	})
	group.POST("/", func(c *gin.Context) {
		ctl.Create(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.PUT("/", func(c *gin.Context) {
		ctl.Update(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.DELETE("/", func(c *gin.Context) {
		ctl.Delete(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	group.DELETE("/destroy", func(c *gin.Context) {
		ctl.Delete(out)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}