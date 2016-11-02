package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	// "github.com/toukii/goutils"
	"html/template"
)

var (
	seq = "AaeHYtVmuIOqKzXl"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs: []template.FuncMap{
			{
				"mod4": mod4,
			},
		},
	}))
	m.Group("/", func() {
		m.Get("", func(ctx *macaron.Context) {
			fmt.Println("word")
			data := make(map[string]interface{})
			data["UserName"] = "shaalx"
			letters := make([]string, 16)
			for i := 0; i < 16; i++ {
				letters[i] = string(seq[i])
			}
			data["seq"] = letters
			ctx.HTML(200, "word", data)
		})
		m.Combo("tube").Get(word)
	})
	m.Run()
}

func word(ctx *macaron.Context) {
	ctx.HTML(200, "tube", nil)
}

func mod4(i int) bool {
	if i == 0 {
		return false
	}
	return i%4 == 0
}
