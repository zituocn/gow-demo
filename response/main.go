package main

import "github.com/zituocn/gow"

func main(){
	r:=gow.Default()

	r.GET("/xml",XMLHandler)

	r.Run()
}


type User struct {

}

func XMLHandler(c *gow.Context){

}