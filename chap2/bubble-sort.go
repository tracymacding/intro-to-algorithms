package main

import ( 
	"math/rand"
	"fmt"
	"time"
)

func bubble_sort_ascending_order(datas []int) {
	for i := 0; i < len(datas); i++ {
		for j := len(datas) - 1; j > i; j-- {
			if datas[j] < datas[j - 1] {
				datas[j], datas[j-1] = datas[j-1], datas[j]
			}
		}
	}
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

	bubble_sort_ascending_order(datas)
	fmt.Printf("After  Ascending sort: %+v\n", datas)	
}
