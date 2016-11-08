package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	"github.com/toukii/wrdGM/util"
	"html/template"
	"time"
)

var (
	sseq = util.SSeq(16)
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
		m.Get("", word)
		m.Post("word", chck)
		m.Combo("tube").Get(tube)
		m.Combo("css").Get(func(ctx *macaron.Context) {
			ctx.HTML(200, "css", nil)
		})
	})
	go ticker()
	macaron.Env = macaron.PROD
	m.Run()
}

func tube(ctx *macaron.Context) {
	ctx.HTML(200, "tube", nil)
}

func word(ctx *macaron.Context) {
	data := make(map[string]interface{})
	data["UserName"] = "shaalx"
	data["seq"] = sseq
	ctx.HTML(200, "word", data)
}

func chck(ctx *macaron.Context) {
	rawPoss := ctx.Query("poss")
	cword := ctx.Query("word")
	word, ok := util.CWord4(rawPoss, sseq)
	if ok && word == cword {
		fmt.Println(word)
		chckRlt := util.ChckCWord(word)
		if chckRlt.Exist {
			ctx.JSON(200, Rslt{Word: word, Score: chckRlt.Score, Code: 200})
		} else {
			fmt.Println("wrong word")
			ctx.JSON(200, Rslt{Code: 301})
		}
	} else {
		fmt.Println(cword, " != ", word)
		ctx.JSON(200, Rslt{Code: 302})
	}
}

type Rslt struct {
	// Word string
	Word  string
	Score int
	Code  int
}

func mod4(i int) bool {
	if i == 0 {
		return false
	}
	return i%4 == 0
}

func ticker() {
	tckr := time.NewTicker(120e9)
	for {
		for {
			sseq1 := util.SSeq(16)
			fmt.Println("sseq:", sseq1)
			rlt := util.SWords(sseq1)
			if rlt.Num >= 35 {
				sseq = sseq1
				break
			}
		}
		<-tckr.C
	}
}
