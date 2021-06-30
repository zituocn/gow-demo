package main

import (
	"github.com/zituocn/gow"
	"time"
)

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	r.Static("/static", "static")
	r.GET("/", htmlHandler)
	r.Run()
}

func htmlHandler(c *gow.Context) {
	c.Data["title"] = "gow html page"
	c.Data["content"] = "基于gow开发的一个网站demo"
	c.Data["date"] = time.Now()
	c.HTML("index.html")
}
