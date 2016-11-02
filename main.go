package main

import "github.com/go-macaron/macaron"
import (
	"fmt"
)

var (
	seq = "AaeHYtVmuIOqKzXl"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{}))

	m.Group("/", func() {
		m.Get("", func(ctx *macaron.Context) {
			fmt.Println("word")
			data := make(map[string]interface{})
			data["UserName"] = "shaalx"
			data["seq"] = seq
			ctx.HTML(200, "word", data)
		})
		m.Combo("tube").Get(word)
	})
	m.Run()
}

func word(ctx *macaron.Context) {
	ctx.HTML(200, "tube", nil)
}
