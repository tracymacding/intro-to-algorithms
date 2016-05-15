package main

import ( 
	"math/rand"
	"fmt"
	"time"
)

func insert_sort_ascending_orde(datas []int) {
	for i := 1; i < len(datas); i++ {
		key := datas[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if key < datas[j] {
				datas[j + 1] = datas[j]
			} else {
				break
			}
		}
		datas[j + 1] = key
	}
}

func insert_sort_descending_orde(datas []int) {
	for i := 1; i < len(datas); i++ {
		key := datas[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if key > datas[j] {
				datas[j + 1] = datas[j]
			} else {
				break
			}
		}
		datas[j + 1] = key
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

	insert_sort_ascending_orde(datas)
	fmt.Printf("After  Ascending sort: %+v\n", datas)	

	insert_sort_descending_orde(datas)
	fmt.Printf("After  Descending sort: %+v\n", datas)	
}
