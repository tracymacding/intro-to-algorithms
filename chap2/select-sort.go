package main

import ( 
	"math/rand"
	"fmt"
	"time"
)

func get_smallest_element_index(datas []int, from, to int) int {
	if datas == nil {
		return -1
	}
	smallest_index := from

	for i := from + 1; i < to; i++ {
		if datas[i] < datas[smallest_index] {
			smallest_index = i
		}
	}
	return smallest_index
} 

func get_largest_element_index(datas []int, from, to int) int {
	if datas == nil {
		return -1
	}
	largest_index := from

	for i := from + 1; i < to; i++ {
		if datas[i] > datas[largest_index] {
			largest_index = i
		}
	}
	return largest_index
} 

func select_sort_ascending_orde(datas []int) {
	for i := 0; i < len(datas); i++ {
		smallest_index := get_smallest_element_index(datas, i, len(datas))
		datas[i], datas[smallest_index] = datas[smallest_index], datas[i]
	}
}

func select_sort_descending_orde(datas []int) {
	for i := 0; i < len(datas); i++ {
		largest_index := get_largest_element_index(datas, i, len(datas))
		datas[i], datas[largest_index] = datas[largest_index], datas[i]
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

	select_sort_ascending_orde(datas)
	fmt.Printf("After  Ascending sort: %+v\n", datas)	

	select_sort_descending_orde(datas)
	fmt.Printf("After  Descending sort: %+v\n", datas)	
}
