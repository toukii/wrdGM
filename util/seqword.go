package util

import (
	"fmt"
	// "github.com/toukii/goutils"
	"strings"
)

var (
	LocMap map[string]int
)

func init() {
	LocMap = map[string]int{
		"0": 0, "1": 1, "2": 2, "3": 3,
		"4": 4, "5": 5, "6": 6, "7": 7,
		"8": 8, "9": 9, "10": 10, "11": 11,
		"12": 12, "13": 13, "14": 14, "15": 15,
	}
}

func SeqWord(possRaw string, letters []string, n int) (word string, ok bool) {
	possRaw = strings.Trim(possRaw, ",")
	poss := strings.Split(possRaw, ",")
	pos := 0
	prePos := 0
	possLen := len(poss)
	// fmt.Println(possLen, possRaw)
	if "" == possRaw || possLen > 16 {
		return
	}
	count := 0
	var exist bool
	for i := 0; i < possLen; i++ {
		it := poss[i]
		if "" == it {
			continue
		}
		exist = false
		pos, exist = LocMap[it]
		if !exist {
			return
		}
		count++
		if count > 1 {
			if !ChckPos(prePos, pos, 4) {
				fmt.Println(i, prePos, pos)
				return
			}
		}
		prePos = pos
		word += letters[pos]
	}
	ok = true
	return
}

func SeqWord4(possRaw string, letters []string) (word string, ok bool) {
	possRaw = strings.Trim(possRaw, ",")
	poss := strings.Split(possRaw, ",")
	pos := 0
	prePos := 0
	possLen := len(poss)
	// fmt.Println(possLen, possRaw)
	if "" == possRaw || possLen > 16 {
		return
	}
	count := 0
	var exist bool
	for i := 0; i < possLen; i++ {
		it := poss[i]
		if "" == it {
			continue
		}
		exist = false
		pos, exist = LocMap[it]
		if !exist {
			return
		}
		count++
		if count > 1 {
			if !ChckPos4(prePos, pos) {
				fmt.Println(i, prePos, pos)
				return
			}
		}
		prePos = pos
		word += letters[pos]
	}
	ok = true
	return
}
