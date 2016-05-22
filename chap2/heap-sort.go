package main

import ( 
	"math/rand"
	"fmt"
	"time"
)

func leftChildIndex(parent int) int {
	return 2*parent + 1
}

func rightChildIndex(parent int) int {
	return 2*parent + 2
}

func maxHeapify(datas []int, rootIndex int) {
	largestIndex := rootIndex
	// 从父节点/左/右孩子节点中选择出最大值
	leftChild := leftChildIndex(rootIndex)
	rightChild := rightChildIndex(rootIndex)
	if leftChild < len(datas) && datas[largestIndex] < datas[leftChild] {
		largestIndex = leftChild
	}
	if rightChild < len(datas) && datas[largestIndex] < datas[rightChild] {
		largestIndex = rightChild
	}
	// 与父节点交换值
	// 并向下递归
	if largestIndex != rootIndex {
		datas[rootIndex], datas[largestIndex] = datas[largestIndex], datas[rootIndex]
		maxHeapify(datas, largestIndex)
	}
} 

// 对数组数据进行堆排序
// 这个算法复杂度为O(n*n)
//func heap_sort_ascending_order(datas []int) {
//	for i := len(datas); i >= 1; i-- {
//		for j := i/2 - 1; j >=0 ; j-- {
//			// datas[i:len(datas)]已经递增有序
//			maxHeapify(datas[0:i], j)
//		}
//		// datas[0]现在是最大值,将其调整至数组末尾
//		// fmt.Printf("datas[0] = %d\n", datas[0])
//		datas[i-1], datas[0] = datas[0], datas[i-1]
//	}
//}

func build_max_heap(datas []int) {
	for i := len(datas)/2 - 1; i >= 0; i-- {
		maxHeapify(datas, i)
	}
}

// 对数组数据进行堆排序
// 这个算法复杂度为O(n*log(n))
func heap_sort_ascending_order(datas []int) {
	// 1. 先建一个大根堆
	build_max_heap(datas)
	// 2. 完成后datas[0]是最大值
	for i := len(datas)-1; i >= 0; i-- {
		datas[i], datas[0] = datas[0], datas[i]
		maxHeapify(datas[0:i], 0)
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
	//k := 1
	//for i := 0; i < 6; i++ {
	//	start := time.Now()
	//	k = k * 10
	//	datas := generate_raw_data(k)
	//	// fmt.Printf("Before sort: %+v\n", datas)	
	//	heap_sort_ascending_order(datas)
	//	cost := time.Since(start).Seconds()
	//	fmt.Printf("Sort %d number, cost %fs\n", k, cost)	
	//}
	// fmt.Printf("After  Ascending sort: %+v\n", datas)	

	datas := generate_raw_data(10)
	fmt.Printf("Before sort: %+v\n", datas)	
	heap_sort_ascending_order(datas)
	fmt.Printf("After  Ascending sort: %+v\n", datas)	
}
