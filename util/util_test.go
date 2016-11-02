package util

import (
	"testing"
)

/*
1  2  3  4
5  6  7  8
9 10 11 12
*/
func TestChckPos(t *testing.T) {
	var ok bool
	// 1
	ok = ChckPos_1(1, 2, 4)
	if !ok {
		t.Error(1, 2, 4)
	}

	ok = ChckPos_1(1, 5, 4)
	if !ok {
		t.Error(1, 5, 4)
	}

	ok = ChckPos_1(1, 6, 4)
	if !ok {
		t.Error(1, 6, 4)
	}
	// 1
	ok = ChckPos_1(2, 1, 4)
	if !ok {
		t.Error(2, 1, 4)
	}

	ok = ChckPos_1(5, 1, 4)
	if !ok {
		t.Error(5, 1, 4)
	}

	ok = ChckPos_1(6, 1, 4)
	if !ok {
		t.Error(6, 1, 4)
	}

	// !1
	ok = ChckPos_1(1, 8, 4)
	if ok {
		t.Error(1, 8, 4)
	}

	ok = ChckPos_1(1, 7, 4)
	if ok {
		t.Error(1, 7, 4)
	}

	ok = ChckPos_1(1, 3, 4)
	if ok {
		t.Error(1, 3, 4)
	}

	ok = ChckPos_1(1, 4, 4)
	if ok {
		t.Error(1, 4, 4)
	}
	// !1
	ok = ChckPos_1(8, 1, 4)
	if ok {
		t.Error(8, 1, 4)
	}

	ok = ChckPos_1(7, 1, 4)
	if ok {
		t.Error(7, 1, 4)
	}

	ok = ChckPos_1(3, 1, 4)
	if ok {
		t.Error(3, 1, 4)
	}

	ok = ChckPos_1(4, 1, 4)
	if ok {
		t.Error(4, 1, 4)
	}

	// 4
	ok = ChckPos_1(4, 8, 4)
	if !ok {
		t.Error(4, 8, 4)
	}
	ok = ChckPos_1(8, 4, 4)
	if !ok {
		t.Error(8, 4, 4)
	}

	ok = ChckPos_1(4, 3, 4)
	if !ok {
		t.Error(4, 3, 4)
	}

	ok = ChckPos_1(3, 4, 4)
	if !ok {
		t.Error(3, 4, 4)
	}

	ok = ChckPos_1(4, 7, 4)
	if !ok {
		t.Error(4, 7, 4)
	}

	ok = ChckPos_1(7, 4, 4)
	if !ok {
		t.Error(7, 4, 4)
	}

	ok = ChckPos_1(4, 8, 4)
	if !ok {
		t.Error(4, 8, 4)
	}

	ok = ChckPos_1(8, 4, 4)
	if !ok {
		t.Error(8, 4, 4)
	}

	// !4
	ok = ChckPos_1(4, 5, 4)
	if ok {
		t.Error(4, 5, 4)
	}

	ok = ChckPos_1(5, 4, 4)
	if ok {
		t.Error(5, 4, 4)
	}

	ok = ChckPos_1(4, 6, 4)
	if ok {
		t.Error(4, 6, 4)
	}

	ok = ChckPos_1(6, 4, 4)
	if ok {
		t.Error(6, 4, 4)
	}

	ok = ChckPos_1(4, 11, 4)
	if ok {
		t.Error(4, 11, 4)
	}

	ok = ChckPos_1(11, 4, 4)
	if ok {
		t.Error(11, 4, 4)
	}

	ok = ChckPos_1(4, 12, 4)
	if ok {
		t.Error(4, 12, 4)
	}

	ok = ChckPos_1(12, 4, 4)
	if ok {
		t.Error(12, 4, 4)
	}

	ok = ChckPos_1(4, 9, 4)
	if ok {
		t.Error(4, 9, 4)
	}

	ok = ChckPos_1(9, 4, 4)
	if ok {
		t.Error(9, 4, 4)
	}

	ok = ChckPos_1(4, 10, 4)
	if ok {
		t.Error(4, 10, 4)
	}

	ok = ChckPos_1(10, 4, 4)
	if ok {
		t.Error(10, 4, 4)
	}
}

func TestChckPos4(t *testing.T) {
	var ok bool
	// 1
	ok = ChckPos4_1(1, 2)
	if !ok {
		t.Error(1, 2, 4)
	}

	ok = ChckPos4_1(1, 5)
	if !ok {
		t.Error(1, 5, 4)
	}

	ok = ChckPos4_1(1, 6)
	if !ok {
		t.Error(1, 6, 4)
	}
	// 1
	ok = ChckPos4_1(2, 1)
	if !ok {
		t.Error(2, 1, 4)
	}

	ok = ChckPos4_1(5, 1)
	if !ok {
		t.Error(5, 1, 4)
	}

	ok = ChckPos4_1(6, 1)
	if !ok {
		t.Error(6, 1, 4)
	}

	// !1
	ok = ChckPos4_1(1, 8)
	if ok {
		t.Error(1, 8, 4)
	}

	ok = ChckPos4_1(1, 7)
	if ok {
		t.Error(1, 7, 4)
	}

	ok = ChckPos4_1(1, 3)
	if ok {
		t.Error(1, 3, 4)
	}

	ok = ChckPos4_1(1, 4)
	if ok {
		t.Error(1, 4, 4)
	}
	// !1
	ok = ChckPos4_1(8, 1)
	if ok {
		t.Error(8, 1, 4)
	}

	ok = ChckPos4_1(7, 1)
	if ok {
		t.Error(7, 1, 4)
	}

	ok = ChckPos4_1(3, 1)
	if ok {
		t.Error(3, 1, 4)
	}

	ok = ChckPos4_1(4, 1)
	if ok {
		t.Error(4, 1, 4)
	}

	// 4
	ok = ChckPos4_1(4, 8)
	if !ok {
		t.Error(4, 8, 4)
	}
	ok = ChckPos4_1(8, 4)
	if !ok {
		t.Error(8, 4, 4)
	}

	ok = ChckPos4_1(4, 3)
	if !ok {
		t.Error(4, 3, 4)
	}

	ok = ChckPos4_1(3, 4)
	if !ok {
		t.Error(3, 4, 4)
	}

	ok = ChckPos4_1(4, 7)
	if !ok {
		t.Error(4, 7, 4)
	}

	ok = ChckPos4_1(7, 4)
	if !ok {
		t.Error(7, 4, 4)
	}

	ok = ChckPos4_1(4, 8)
	if !ok {
		t.Error(4, 8, 4)
	}

	ok = ChckPos4_1(8, 4)
	if !ok {
		t.Error(8, 4, 4)
	}

	// !4
	ok = ChckPos4_1(4, 5)
	if ok {
		t.Error(4, 5, 4)
	}

	ok = ChckPos4_1(5, 4)
	if ok {
		t.Error(5, 4, 4)
	}

	ok = ChckPos4_1(4, 6)
	if ok {
		t.Error(4, 6, 4)
	}

	ok = ChckPos4_1(6, 4)
	if ok {
		t.Error(6, 4, 4)
	}

	ok = ChckPos4_1(4, 11)
	if ok {
		t.Error(4, 11, 4)
	}

	ok = ChckPos4_1(11, 4)
	if ok {
		t.Error(11, 4, 4)
	}

	ok = ChckPos4_1(4, 12)
	if ok {
		t.Error(4, 12, 4)
	}

	ok = ChckPos4_1(12, 4)
	if ok {
		t.Error(12, 4, 4)
	}

	ok = ChckPos4_1(4, 9)
	if ok {
		t.Error(4, 9, 4)
	}

	ok = ChckPos4_1(9, 4)
	if ok {
		t.Error(9, 4, 4)
	}

	ok = ChckPos4_1(4, 10)
	if ok {
		t.Error(4, 10, 4)
	}

	ok = ChckPos4_1(10, 4)
	if ok {
		t.Error(10, 4, 4)
	}

}

func BenchmarkChckPos_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChckPos_1(0, 3, 4)
		ChckPos_1(7, 0, 4)
	}
}

func BenchmarkChckPos4_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChckPos4_1(0, 3)
		ChckPos4_1(7, 0)
	}
}

func BenchmarkChckPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChckPos(0, 3, 4)
		ChckPos(7, 0, 4)
	}
}

func BenchmarkChckPos4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChckPos4(0, 3)
		ChckPos4(7, 0)
	}
}
