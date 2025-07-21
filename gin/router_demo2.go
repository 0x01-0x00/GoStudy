package gin

import (
	"GoStudy/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterDemo2() {
	// 模板路由
	groupUser := global.R.Group("/tmpl")
	groupUser.GET("/index", func(c *gin.Context) {
		// 加载模板
		global.R.LoadHTMLGlob("/Users/zhongtao/GolandProjects/GoStudy/tmpl/index.tmpl")
		// 渲染模板
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
		})
	})
	groupUser.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	groupUser.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	groupUser.DELETE("/index", func(context *gin.Context) {

	})
}
