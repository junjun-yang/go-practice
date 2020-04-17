package main

import (
	"fmt"
)

/**
测试发送和接收数据
*/
func testSendAndReceived() {
	c := make(chan int)
	//defer close(c)
	go func() { c <- 3 + 4 }()
	i := <-c
	fmt.Println("expect received 7 and actual received ", i)
	close(c)
	x, ok := <-c
	//从一个已经关闭的channel中接收数据，会返回默认值和false
	fmt.Println("expect received (0,false) and actual received ", x, ok)
}

/**
利用管道阻塞，实现同步,多个线程向channel中写入，并不保证写入顺序
*/
func testBlock() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	//time.Sleep(time.Duration(2)*time.Second)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
func sum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum
}
