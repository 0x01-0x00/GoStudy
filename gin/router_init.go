package gin

import (
	"GoStudy/global"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	var err error
	// 创建路由
	global.R = gin.Default()

	// 测试路由
	global.R.GET("/gin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	RouterDemo1()
	RouterDemo2()

	err = global.R.Run(":8080")
	if err != nil {
		return
	}
}
