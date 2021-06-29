package main

import (
	"github.com/zituocn/gow"
)

func main() {
	r := gow.Default()

	v1 := r.Group("/v1")

	// /v1 下所有路由，使用APIAuth()
	v1.Use(APIAuth())

	// /v1/data
	v1.GET("/data", func(c *gow.Context) {
		c.JSON(gow.H{
			"code": 0,
			"msg":  "OK",
		})
	})

	user := v1.Group("/user")

	// /v1/user/ 下所有路由，使用UserAuth()
	user.Use(UserAuth())
	{
		// route: /v1/user/1
		user.GET("/{uid}", UserHandler)
	}
	r.Run()

}

// APIAuth API接口鉴权
func APIAuth() gow.HandlerFunc {
	return func(c *gow.Context) {
		auth := c.GetHeader("auth")
		if auth != "123456" {
			c.ServerJSON(403, gow.H{
				"code": 403,
				"msg":  "没有API访问权限",
			})
			c.StopRun()
			return
		}
		c.Next()
	}
}

// UserAuth 用户鉴权
func UserAuth() gow.HandlerFunc {
	return func(c *gow.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.ServerJSON(403, gow.H{
				"code": 403,
				"msg":  "此用户无权限",
			})
			c.StopRun()
			return
		}
		c.Next()
	}
}

// UserHandler get user info
func UserHandler(c *gow.Context) {
	h := map[string]interface{}{
		"uid":      1,
		"username": "zituocn",
		"city":     "成都市",
		"prov":     "四川省",
	}
	c.JSON(h)
}
