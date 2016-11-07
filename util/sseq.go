package util

/*
按照字母的概率，累加概率和，方便定位概率对应的字母
*/
import (
	"math/rand"
	"sort"
	"time"
)

var (
	percent = map[string]int{
		"e": 1268, "t": 978, "a": 788, "o": 776, "i": 707,
		"n": 706, "s": 634, "r": 594, "h": 573, "l": 394,
		"d": 389, "u": 280, "c": 268, "f": 256, "m": 244,
		"w": 214, "y": 202, "g": 187, "p": 186, "b": 156,
		"v": 102, "k": 60, "x": 16, "j": 10, "q": 9, "z": 6,
		"er": 122, "in": 98, "on": 74, "ti": 71, "at": 57,
		"re": 57, "te": 54, "en": 50, "le": 32, "an": 30,
	}
	seqs     = []string{"e", "t", "a", "o", "i", "n", "s", "r", "h", "l", "d", "u", "c", "f", "m", "w", "y", "g", "p", "b", "v", "k", "x", "j", "q", "z", "er", "in", "on", "ti", "at", "re", "te", "en", "le", "an"}
	perSlice []int
	locSeq   []string

	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	perSlice = make([]int, 0, 36)
	locSeq = make([]string, 0, 36)
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
func SSeq(sseqLen int) []string {
	filter := make(map[string]struct{})
	sseqs := make([]string, 0, sseqLen)
	n := perSlice[len(perSlice)-1]
	for i := 0; i < sseqLen; i++ {
		per := rander.Intn(n)
		seq := searchSeq(per)
		if _, ok := filter[seq]; ok {
			i--
			continue
		}
		filter[seq] = struct{}{}
		if seq != "" {
			sseqs = append(sseqs, seq)
		}
	}
	return sseqs
}
