package main

import "github.com/zituocn/gow"

func main() {

	r := gow.Default()

	// GET /
	r.GET("/", IndexHandler)

	// GET /user/1
	r.GET("/user/1", UserHandler)

	// POST /user/1
	r.POST("/user/1", UserUpdateHandler)

	// DELETE /user/1
	r.DELETE("/user/1", UserDeleteHandler)

	// PUT /user/1
	r.PUT("/user/1", UserPutHandler)

	// /article/100
	r.GET("/article/{article_id}", ArticleDetailHandler)

	// /read-100.html
	r.GET("/read-{id}.html", ReadHandler)

	// /read/test/101
	r.GET("/{name}/{id}", NameHandler)

	// route group
	v1 := r.Group("/v1")
	{
		// /v1/user/1
		v1.GET("/user/1", UserHandler)

		// /v1/user/1
		v1.POST("/user/1", UserUpdateHandler)
	}

	// 静态资源路由
	r.Static("/static", "static")

	r.Run(8080)
}

func IndexHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func UserHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func UserUpdateHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func UserDeleteHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func UserPutHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func ArticleDetailHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func ReadHandler(c *gow.Context) {
	c.String(c.Request.RequestURI)
}

func NameHandler(c *gow.Context) {
	name := c.Param("name")
	id, _ := c.ParamInt64("id")
	c.JSON(gow.H{
		"name": name,
		"id":   id,
	})
}
