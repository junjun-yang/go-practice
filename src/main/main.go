package main

import (
	"fmt"
	"math"
	"time"
)

/**
简单函数
*/
func add(x, y int) int {
	return x + y
}

/**
多个返回值函数
*/
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum + x
	return
}

/**
条件语句
*/
func testIf(cause int, result int) int {
	//可以在if语句内定义局部变量
	if x := math.Pow(10, float64(cause)); x < 10000 {
		println("success", x)
	} else {
		println("failed", x)
	}
	return result
}

/**
循环函数
*/
func testLoop() {
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
}

/**
延迟函数，被Defer关键字修饰的方法会被先压到栈中，延迟执行，待上层方法全部执行完之后才被调用
典型的场景可用于多层输入输出流的关闭
*/
func testDefer() {
	fmt.Println("i am first")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("i am second")
}

/**
测试switch语句：不需要break语句，遇到匹配的case会立即停止，除非显式声明fallthrough语句
*/
func testSwitch(x int) {
	switch x {
	case 0:
		println("case 0")
	case 1:
		println("case 1")
	}
	switch x {
	case 0:
		println("fallthrough case 0")
		fallthrough
	case 1:
		println("fallthrough case 1")
		fallthrough
	default:
		println("fallthrough default")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("time hour < 12", t)
	case t.Hour() > 15:
		fmt.Println("time hour > 15", t)
	}
}

/**
测试指针，指针中保存的是目标对象的内存地址 *指针代表真实的目标对象,可以直接通过修改*ptr改变目标对象的值
*/
func testPtr() {
	var ptr *int
	i := 12
	ptr = &i
	fmt.Println("ptr==", ptr)
	fmt.Println("*ptr==", *ptr)
	fmt.Println("i==", i)
	*ptr = 21
	fmt.Println("*ptr assign to 21")
	fmt.Println("*ptr==", *ptr)
	fmt.Println("i==", i)
	fmt.Println("ptr's address ", &ptr)
}

type Vertex struct {
	X int
	Y int
}

/**
测试结构体
*/
func testStruct() {
	x := Vertex{1, 2}
	y := Vertex{X: 1}
	z := Vertex{}
	ptr := &Vertex{3, 4}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(*ptr)
}


/**
测试切片（可动态改变大小的数组，可理解为动态数组）
有两种创建切片的方式
*/
func testSlice() {
	slice := make([]int, 0, 10)
	fmt.Println(len(slice), cap(slice))
	slice = []int{1, 2, 3, 4}
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice[:2])
	fmt.Println(slice[0:4])
	fmt.Println(append(slice, 5, 6)[0:])
	for index, value := range slice {
		fmt.Println(index, value)
	}
}

/**
测试map
map赋值：	mymap["aa"]=1
检测map是否存在：	elem,ok:=mymap["aa"]，ok为true存在，ok为false不存在
删除键值对：  delete(mymap,"aa")
*/
func testMap() {
	mymap := make(map[string]int)
	mymap["aa"] = 1
	mymap["bb"] = 2
	println(mymap["aa"])
	println(mymap["cc"])
	elem, ok := mymap["cc"]
	fmt.Println(elem, ok)
	delete(mymap, "bb")
	fmt.Println(len(mymap))
}

/**
测试函数闭包,函数可以看做一个普通的变量，我们可以直接对它进行访问和赋值
*/
func testClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

/**
返回值为一个函数
*/
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

/**
go中定义之后未使用的变量，编译时会报错
*/
func main() {
	println(fmt.Sprintf("YT%d",time.Now().UnixNano()/1e6))
	//testSendAndReceived()
	//testBlock()
	//testPtr()
}
