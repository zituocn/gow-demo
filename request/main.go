package main

import (
	"github.com/zituocn/gow"
)

func main() {
	r := gow.Default()
	r.GET("/v1/user", GetUser)
	r.PUT("/v1/user", UpdateUser)
	r.Run()
}

// GetUser get user
func GetUser(c *gow.Context) {
	name := c.GetString("name")
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt64("limit")
	is, _ := c.GetBool("is")
	score, _ := c.GetFloat64("score")

	h := gow.H{
		"name":  name,
		"page":  page,
		"limit": limit,
		"is":    is,
		"score": score,
	}
	c.JSON(h)
}

type UserInfo struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

// UpdateUser 更新用户
func UpdateUser(c *gow.Context) {
	userInfo := new(UserInfo)

	// 反序列化到 struct
	err := c.DecodeJSONBody(&userInfo)
	if err != nil {
		c.JSON(gow.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	// 输出
	c.JSON(userInfo)
}
