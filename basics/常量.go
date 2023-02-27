package main

import (
	"fmt"
	"math"
)

// 一： 常量
func conSts() {
	const (
		filename = "笔记.txt" //常量也可以放在函数外面，为包内的函数提供调用。
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 二：枚举类型
func meiJu() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

// 三：iota 参与运算
func meiJu2() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
