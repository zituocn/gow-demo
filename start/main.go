package main

import (
	"github.com/zituocn/gow"
)

func main() {
	r := gow.Default()
	r.GET("/", func(c *gow.Context) {
		c.JSON(gow.H{
			"code": 0,
			"msg":  "OK",
		})
	})

	r.GET("/hello", hello)

	//default :8080
	r.Run()
}

func hello(c *gow.Context) {
	c.String("hello gow!")
}
