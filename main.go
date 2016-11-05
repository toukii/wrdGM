package main

import (
	"fmt"
	"github.com/go-macaron/macaron"
	"github.com/toukii/wrdGM/util"
	"html/template"
	"time"
)

var (
	sseq    = util.SSeq(16)
	letters = make([]string, 16)
	// swords *util.Resultwords
)

func init() {
	for i := 0; i < 16; i++ {
		letters[i] = string(sseq[i])
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
	go ticker()
	m.Run()
}

func tube(ctx *macaron.Context) {
	ctx.HTML(200, "tube", nil)
}

func word(ctx *macaron.Context) {
	data := make(map[string]interface{})
	data["UserName"] = "shaalx"
	data["seq"] = letters
	ctx.HTML(200, "word", data)
}

func chck(ctx *macaron.Context) {
	rawPoss := ctx.Query("poss")
	cword := ctx.Query("word")
	word, ok := util.CWord4(rawPoss, letters)
	if ok && word == cword {
		fmt.Println(word)
		chckRlt:=util.ChckCWord(word)
		if chckRlt.Exist {
			ctx.JSON(200, Rslt{Word: word,Score:chckRlt.Score, Code: 200})	
		}else{
			fmt.Println("wrong word")
			ctx.JSON(200, Rslt{ Code: 301})
		}
	} else {
		fmt.Println(cword, " != ", word)
		ctx.JSON(200, Rslt{ Code: 302})
	}
}

type Rslt struct {
	// Word string
	Word string
	Score int
	Code int
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
		fmt.Println("sseq:", sseq)
		_ = util.SWords(sseq)
		for i := 0; i < 16; i++ {
			letters[i] = string(sseq[i])
		}
		<-tckr.C
		sseq = util.SSeq(16)
	}
}
