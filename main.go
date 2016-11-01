package main

import "github.com/go-macaron/macaron"
import (
	"fmt"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		// Layout:".html",
	}))
	m.Get("/",func (ctx *macaron.Context)  {
		fmt.Println("word")
		data:=make(map[string]interface{})
		data["UserName"]="shaalx"
		ctx.HTML(200,"word",data)
	})
	// m.Group("/",func ()  {
	// 	fmt.Println("/hello")

	// 	m.Combo("/v1").Get(word)
	// })
	m.Run()
}

func word(ctx *macaron.Context)  {
	// ctx.HTML(200,"word.html",nil)
	ctx.HTMLSet(200,"word","word.html",nil)
}