package util

/*
判断两个方块是否相连
*/

// 下标从0开始
func ChckPos(i, j, rowNum int) bool {
	// 在上下挨着，或一行上挨着
	if (i%rowNum == j%rowNum && abs(i/rowNum, j/rowNum) == 1) || (i/rowNum == j/rowNum && abs(i, j) == 1) {
		return true
	}
	// 斜对角
	if abs(i/rowNum, j/rowNum) == 1 && abs(i%rowNum, j%rowNum) == 1 {
		return true
	}
	return false
}

// 下标从1开始
func ChckPos_1(i, j, rowNum int) bool {
	i, j = i-1, j-1
	return ChckPos(i, j, rowNum)
}

// 下标从0开始
func ChckPos4(i, j int) bool {
	// 在上下挨着，或一行上挨着
	if (i%4 == j%4 && abs(i>>2, j>>2) == 1) || (i>>2 == j>>2 && abs(i, j) == 1) {
		return true
	}
	// 斜对角
	if abs(i>>2, j>>2) == 1 && abs(i%4, j%4) == 1 {
		return true
	}
	return false
}

// 下标从1开始
func ChckPos4_1(i, j int) bool {
	i, j = i-1, j-1
	return ChckPos4(i, j)
}

func abs(i, j int) int {
	v := i - j
	if v > 0 {
		return v
	}
	return -v
}
