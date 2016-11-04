package util

/*
按照字母的概率，累加概率和，方便定位概率对应的字母
*/
import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

var (
	percent = map[string]int{
		"e": 1268, "t": 978, "a": 788, "o": 776, "i": 707,
		"n": 706, "s": 634, "r": 594, "h": 573, "l": 394,
		"d": 389, "u": 280, "c": 268, "f": 256, "m": 244,
		"w": 214, "y": 202, "g": 187, "p": 186, "b": 156,
		"v": 102, "k": 60, "x": 16, "j": 10, "q": 9, "z": 6,
	}
	seqs     = []string{"e", "t", "a", "o", "i", "n", "s", "r", "h", "l", "d", "u", "c", "f", "m", "w", "y", "g", "p", "b", "v", "k", "x", "j", "q", "z"}
	perSlice []int
	locSeq   []string

	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	perSlice = make([]int, 0, 26)
	locSeq = make([]string, 0, 26)
	perSum := 0
	seqsLen := len(seqs)
	for i := 0; i < seqsLen; i++ {
		seq := seqs[i]
		per := percent[seq]
		perSum += per
		perSlice = append(perSlice, perSum)
		locSeq = append(locSeq, seq)
	}

}

// 根据概率，定位字母
/*BenchmarkSearchSeq2-4           50000000                35.7 ns/op*/
func searchSeq2(v int) string {
	// for i, per := range perSlice {
	// 	if v <= per {
	// 		return locSeq[i]
	// 	}
	// }
	loc := sort.SearchInts(perSlice, v)
	return locSeq[loc]
}

/*BenchmarkSearchSeq-4            500000000                3.54 ns/op*/
func searchSeq(v int) string {
	for i, per := range perSlice {
		if v <= per {
			return locSeq[i]
		}
	}
	return ""
}

// 随机数数组:每一项对应一个字母序列的索引(查找得到)
func SSeq() string {
	sseqs := make([]string, 0, 16)
	n := perSlice[25]
	for i := 0; i < 16; i++ {
		per := rander.Intn(n)
		seq := searchSeq(per)
		if seq != "" {
			sseqs = append(sseqs, seq)
		}
	}
	return strings.Join(sseqs, "")
}
