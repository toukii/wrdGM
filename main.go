package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	"github.com/toukii/goutils"
	"github.com/toukii/wrdGM/util"
	"html/template"
	"strconv"
	"strings"
)

var (
	seq        = "AaeHYtVmuIOqKzXl"
	letters    = make([]string, 16)
	lettersMap = make(map[string]string)
)

func init() {
	for i := 0; i < 16; i++ {
		letters[i] = string(seq[i])
		lettersMap[fmt.Sprintf("%d", i)] = letters[i]
	}
}

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
		m.Get("", word)
		m.Post("word", chck)
		m.Combo("tube").Get(tube)
	})
	m.Run()
}

func tube(ctx *macaron.Context) {
	ctx.HTML(200, "tube", nil)
}

func word(ctx *macaron.Context) {
	fmt.Println("word")
	data := make(map[string]interface{})
	data["UserName"] = "shaalx"
	data["seq"] = letters
	ctx.HTML(200, "word", data)
}

func chck(ctx *macaron.Context) {
	possRaw := ctx.Query("poss")
	poss := strings.Split(possRaw, ",")
	words := ""
	for i := 0; i < len(poss)-1; i++ {
		it := poss[i]
		pos, err := strconv.ParseInt(it, 10, 16)
		if goutils.CheckErr(err) {
			ctx.JSON(302, err.Error())
		}
		if i > 0 {
			if !util.ChckPos4(i-1, i) {
				ctx.JSON(303, "pos error")
			}
		}
		words += letters[pos]
	}
	fmt.Println(words)
	ctx.JSON(200, words)
}

func mod4(i int) bool {
	if i == 0 {
		return false
	}
	return i%4 == 0
}
