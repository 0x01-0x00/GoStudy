package gin

import (
	"GoStudy/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterDemo1() {
	// 定义路由
	global.R.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	global.R.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	global.R.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	global.R.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
}
