package gin

import (
	"GoStudy/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Apple struct {
	Iphone  string
	Mac     string
	Watch   string
	Airpods string
	ipad    string // 小写私有字段，不会被 json 包序列化
	Ipod    string `json:"ipod"`
}

func RouterDemo1() {
	apple := Apple{
		Iphone:  "Iphone",
		Mac:     "Mac",
		Watch:   "Watch",
		Airpods: "Airpods",
		ipad:    "ipad", // 小写私有字段，不会被 json 包序列化
		Ipod:    "ipod",
	}
	// 定义路由
	global.R.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "GET",
			"RouterDemo1": apple,
		})
	})
	global.R.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "POST",
			"RouterDemo1": apple,
		})
	})
	global.R.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "PUT",
			"RouterDemo1": apple,
		})
	})
	global.R.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "DELETE",
			"RouterDemo1": apple,
		})
	})
}
