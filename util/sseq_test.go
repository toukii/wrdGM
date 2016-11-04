package util

import (
	"testing"
)

var (
	pers = map[string]int{
		"e": 1268, "t": 2246, "a": 3034, "o": 3810, "i": 4517,
		"n": 5223, "s": 5857, "r": 6451, "h": 7024, "l": 7418,
		"d": 7807, "u": 8087, "c": 8355, "f": 8611, "m": 8855,
		"w": 9069, "y": 9271, "g": 9458, "p": 9644, "b": 9800,
		"v": 9902, "k": 9962, "x": 9978, "j": 9988, "q": 9997, "z": 10003,
	}
)

func TestInit(t *testing.T) {
	for seq, per := range pers {
		seq_search := searchSeq(per)
		if seq_search != seq {
			t.Log(seq, " != ", seq_search)
		}
		seq_search2 := searchSeq(per - 5)
		if seq_search2 != seq {
			t.Log(seq, " != ", seq_search)
		}
	}
}

func BenchmarkSearchSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchSeq(888)
	}
}

func BenchmarkSearchSeq2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchSeq2(888)
	}
}

func TestSSeq(t *testing.T) {
	sseq := SSeq()
	t.Log(sseq)
}
