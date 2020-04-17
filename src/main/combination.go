package main

import (
	"fmt"
	"math"
)

/**
圆
*/
type Circle struct {
	x, y, r float64
}

/**
go中结构体的方法需要在外部定义
*/
func (this *Circle) Area() float64 {
	return math.Pi / 2 * this.r * this.r
}

type Rect struct {
	x, y, w, h float64
}

func (this *Rect) Area() float64 {
	return this.w * this.h
}

/**
求面积之和
此方法传入的参数为任意实现了Area()方法的结构体集合,相比于Java，抛弃了继承和实现的概念
*/
func Area(shapes ...interface{ Area() float64 }) float64 {
	var result float64
	for _, shape := range shapes {
		result += shape.Area()
	}
	return result
}
/**
测试组合
 */
func testArea()  {
	area := Area(&Circle{r: 20}, &Rect{w: 10, h: 10})
	fmt.Println(area)
}
