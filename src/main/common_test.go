package main

import (
	"fmt"
	"testing"
)

/**
测试数组
*/
func TestArray(t *testing.T) {
	array := [3]int{1, 2, 3}
	array2 := [...]int{2: 3, 3: 3}
	println(len(array))
	println(cap(array))
	for i := range array {
		println(array[i])
	}
	for i, i2 := range array2 {
		fmt.Println(i, i2)
	}
}
