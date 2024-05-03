package main

import (
	"fmt"
)

func main() {
	// 多次元連想配列は都度その次元でmapを作成する必要がある
	arr := make(map[int]map[int]int)
	arr[0] = make(map[int]int)
	arr[0][0] = 0
	fmt.Println(arr)

	m := map[int]map[int]int{}
	var cnt int
	for i := 0; i < 5; i++ {
		m[i] = map[int]int{}
		for k := 0; k < 5; k++ {
			cnt ++
			m[i][k] = cnt
		}
	}
	fmt.Println(m)
}
