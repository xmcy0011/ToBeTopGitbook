package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 导出切片结构示例
func exportSlice() {
	var s1 = []int{1, 2, 3}
	s2 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("%v, 底层数组地址: %d, Len: %d, Cap: %d", s2, s2.Data, s2.Len, s2.Cap)
}

// nil切片和空切片的区别
func nilSliceAndZeroSlice() {
	var s1 []int
	s2 := make([]int, 0)
	s3 := make([]int, 0)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s3)))
}

// 数组初始化2中方式
func array() {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	arr2 := [2]int{1, 2}
	arr2 = [...]int{1, 2}

	fmt.Println("is contains 3?", contains(arr2, 3))
}

// 数组长度是类型的组成部分
func arrayType() {
	var arr1 [2]int
	// var arr2 [3]int

	contains(arr1, 1)
	// contains(arr2, 1) // 报错：Cannot use 'arr2' (type [3]int) as the type [2]int
}

func contains(arr [2]int, num int) bool {
	for i := range arr {
		if i == num {
			return true
		}
	}
	return false
}

func growSlice() {
	var s []int
	s = append(s, 0, 1, 2)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 4)
	fmt.Println(s, len(s), cap(s))
}

func main() {
	growSlice()
}
