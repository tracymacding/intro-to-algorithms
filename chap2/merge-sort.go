package main

import ( 
	"math/rand"
	"fmt"
	"time"
)

const (
	MaxInt64 = 1 << 63 - 1 
)

// datas[p ... q]与datas[q + 1 ... r]分别有序
// 现在将其合并成整体有序
func merge(datas []int, p, q, r int) {

	s1 := make([]int, 0, q-p+2)
	s2 := make([]int, 0, r-q+1)

	for i := p; i < q + 1 ; i++ {
		s1 = append(s1, datas[i])
	}
	s1 = append(s1, MaxInt64)

	for i := q+1; i < r+1; i++ {
		s2 = append(s2, datas[i])
	}
	s2 = append(s2, MaxInt64)

	// fmt.Printf("Start merge %+v and %+v\n", datas[p:q + 1], datas[q+1:r + 1])

	i, j := 0, 0
	for k := p; k < r+1; k++ {
		if s1[i] <= s2[j] {
			datas[k] = s1[i]
			i++
		} else {
			datas[k] = s2[j]
			j++
		}
	}

	// fmt.Printf("merge result: %+v\n", datas[p:r + 1])
}

func merge_sort_ascending_order(datas []int, start, end int) {
	if (start >= end) {
		return
	}

	middle := (start + end) / 2
	//fmt.Printf("Start merge sort %+v\n", datas[start:end+1])
	merge_sort_ascending_order(datas, start, middle)
	merge_sort_ascending_order(datas, middle + 1, end)
	
	merge(datas, start, middle, end)
}

func generate_raw_data(n int) []int {
	res := make([]int, 0, n)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		res = append(res, r.Intn(100))
	}
	return res
}

func main() {
	datas := generate_raw_data(10)
	fmt.Printf("Before sort: %+v\n", datas)	

	merge_sort_ascending_order(datas, 0, len(datas) - 1)
	fmt.Printf("After  Ascending sort: %+v\n", datas)
}
