package util

import (
	"fmt"
	"testing"
)

/*
AaeH
YtVm
uIOq
KzXl
*/

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

func TestSeqWord(t *testing.T) {
	possRaw := ""
	word, ok := SeqWord(possRaw, letters, 4)
	if ok {
		t.Error(possRaw, word, ok)
	}

	possRaw = "0,1,2,3,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = "0,1,2,3,,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = "0,1,2,,3,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = ",0,1,2,,3,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = ",0,1,2,3,7,6,5,4,8,9,10,11,15,14,13,12,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeHmVtYuIOqlXzK" != word {
		t.Error(possRaw, word, ok)
	}

	possRaw = ",0,1,2,3,7,11,15,14,13,12,8,4,5,10,6,9,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AaeHmqlXzKuYtOVI" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = ",0,5,10,11,15,14,13,9,12,8,4,1,6,3,7,2,"
	word, ok = SeqWord(possRaw, letters, 4)
	if !ok || "AtOqlXzIKuYaVHme" != word {
		t.Error(possRaw, word, ok)
	}

}

func BenchmarkSeqWord(b *testing.B) {
	possRaw := ",0,5,10,11,15,14,13,9,12,8,4,1,6,3,7,2,"
	for i := 0; i < b.N; i++ {
		SeqWord(possRaw, letters, 4)
	}
}

func BenchmarkSeqWord4(b *testing.B) {
	possRaw := ",0,5,10,11,15,14,13,9,12,8,4,1,6,3,7,2,"
	for i := 0; i < b.N; i++ {
		SeqWord4(possRaw, letters)
	}
}
