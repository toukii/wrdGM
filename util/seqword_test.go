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
	word, ok := SeqWord(possRaw, letters)
	if ok {
		t.Error(possRaw, word, ok)
	}

	possRaw = "0,1,2,3,"
	word, ok = SeqWord(possRaw, letters)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = "0,1,2,3,,"
	word, ok = SeqWord(possRaw, letters)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = "0,1,2,,3,"
	word, ok = SeqWord(possRaw, letters)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
	possRaw = ",0,1,2,,3,"
	word, ok = SeqWord(possRaw, letters)
	if !ok || "AaeH" != word {
		t.Error(possRaw, word, ok)
	}
}
