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

	// GET /:param1/:param2
	global.R.GET("/url/:param1/:param2", func(c *gin.Context) {
		param1 := c.Param("param1")
		param2 := c.Param("param2")
		c.JSON(http.StatusOK, gin.H{
			"message":     "GET",
			"RouterDemo1": apple,
			"param1":      param1,
			"param2":      param2,
		})
	})

	// GET /user
	global.R.GET("/user", func(c *gin.Context) {
		//name := c.Query("name") // 获取 query 参数
		name := c.DefaultQuery("name", "0x00") // 获取 query 参数，如果没有则返回默认值
		c.JSON(http.StatusOK, gin.H{
			"message":     "GET",
			"RouterDemo1": apple,
			"name":        name,
		})
	})
	// POST /user
	global.R.POST("/user", func(c *gin.Context) {
		user := c.PostForm("user")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"message":     "POST",
			"RouterDemo1": apple,
			"user":        user,
			"password":    password,
		})
	})
	// PUT /user
	global.R.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "PUT",
			"RouterDemo1": apple,
		})
	})
	// DELETE /user
	global.R.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "DELETE",
			"RouterDemo1": apple,
		})
	})
}
