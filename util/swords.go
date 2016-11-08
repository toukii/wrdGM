package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// 无向图，4*4=16个节点。 即在这个序列图上生成字母组合
type Graph [4][4]Neighbour

// 对应无向图节点是否访问过
type Vis [4][4]bool

// 无向图节点，
// V：该节点的值（字母）
// E：邻接点（相邻的节点）
type Neighbour struct {
	V string
	E []Edge
}

// 邻接点信息，I行 J列
type Edge struct {
	I, J int
}

// 单词词库
type Ciku map[string]bool

// 全局变量
var (
	N           int         //N行N列（N=4）
	vis         Vis         // 记录是否访问过
	ciku        *Ciku       //单词库，预加载
	result      chan string // 生成单词后放入该通道
	resultwords Resultwords //结果
	graph       *Graph      // 无向图
)

type Words map[string]int

type Resultwords struct { //玩家游戏结果数据
	Sequence []string //字母序列
	Score    int      //得分
	Num      int      //单词个数
	ID       string   //ID
	Words    *Words   //找到的单词
}

// 将无向图写出到out(.json) 中
func (t Graph) write(out string) {
	file, _ := os.Create(out)
	defer file.Close()
	b, _ := json.Marshal(t)
	ioutil.WriteFile(out, b, 0644)
}

// 从in(.json)读取无向图
func read(in string) *Graph {
	b, _ := ioutil.ReadFile(in)
	var v Graph
	json.Unmarshal(b, &v)
	return &v
}

// 从in(.json)初始化无向图：
func initGraphFromJson(in string) *Graph {
	t := read(in)
	return t
}

// 使用字母序列初始化无向图
func (t *Graph) inti_graph_with_sequence(sequence []string) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			index := i*N + j
			t[i][j].V = strings.ToLower(string(sequence[index]))
		}
	}
}

// 词库
func FromJson(in string) (*Ciku, int) {
	b, _ := ioutil.ReadFile(in)
	var ciku Ciku
	json.Unmarshal(b, &ciku)
	return &ciku, len(ciku)
}

// 初始化全局变量
func initAll() {
	N = 4
	vis = *initVis()
	result = make(chan string, 1)
	ciku, _ = FromJson("word_json.json")
	graph = initGraphFromJson("graph.json")
}

// 初始化访问标志   （其实没有必要）
func initVis() *Vis {
	return &vis
}

// 核心代码，深度优先算法遍历该无向图，若单词合法，将其压入单词通道result中
// e：当前访问的位置
// deep:访问的深度，调节deep的值可以改变整体性能
// s：当前访问得到的字母组合
func (t Graph) dfs(e Edge, deep int, s string) {
	if vis[e.I][e.J] || deep > 15 {
		return
	}
	s = s + t[e.I][e.J].V
	if (*ciku)[s] {
		result <- s
	}
	vis[e.I][e.J] = true
	for _, item := range t[e.I][e.J].E {
		t.dfs(item, deep+1, s)
	}
	vis[e.I][e.J] = false
}

// 读取result单词通道，最后返回得到的单词
// q ：结束命令，若q中有数据，表示可以return了
func readChan(q chan struct{}) *(Words) {
	m := make(Words)
	quit := false
	for {
		select {
		case <-q:
			quit = true
		case r := <-result:
			fmt.Printf("%16s", r)
			m[r] = 10 * len(r)
		}
		if quit {
			break
		}
	}
	resultwords.Words = &m
	resultwords.Num = len(m)
	if resultwords.Num > 25 {
		// resultwords.resultWords_to_json()
	}
	return &m
}

// 将结果out(map[string]bool)输出到out(.json)
func resultToJson(out string, m *map[string]bool) {
	of, _ := os.Create(out)
	defer of.Close()
	b, _ := json.Marshal(m)
	of.Write(b)
}

// 读取结果：测试用（可以删除）
func resultFromJson(in string) {
	var m map[string]bool
	b, _ := ioutil.ReadFile(in)
	json.Unmarshal(b, &m)
	for item := range m {
		fmt.Println(item)
	}
}

// 开始
func start(sequence []string) {
	resultwords.Sequence = sequence
	graph.inti_graph_with_sequence(sequence)
	one_finding(graph)
}

var (
	t1, t2 time.Time
	qb     chan struct{}
)

// 做一次序列的查找
func one_finding(graph *Graph) {
	var e Edge
	var s string
	qb = make(chan struct{}, 1)
	t1 = time.Now()
	var m Words
	go func(m *Words) {
		m = readChan(qb)
		fmt.Printf("\nEND, %d word found , ", len(*m))
	}(&m) // go readChan(qb)
	fmt.Println("start...")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			e.I = i
			e.J = j
			s = ""
			graph.dfs(e, 0, s)
		}
	}
	time.Sleep(1)
	qb <- struct{}{}
	time.Sleep(1)
	t2 = time.Now()
	fmt.Println(t2.Sub(t1))
}

func SWords(sseq []string) *Resultwords {
	initAll()
	start(sseq)
	return &resultwords
}

type ChckRlt struct {
	Exist bool
	Score int
}

func ChckCWord(cword string) (chckRlt ChckRlt) {
	if resultwords.Words == nil {
		return
	}
	score, exist := (*resultwords.Words)[strings.ToLower(cword)]
	if exist {
		chckRlt.Exist = true
		chckRlt.Score = score
	}
	return
}
