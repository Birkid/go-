package main

import "fmt"

// 一： 给变量定义类型
func variable() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

// 二： 给变量赋值
func fuZhi() {
	//var a int = 3				//给 a 定义为 int 并赋值
	var a, s int = 3, 4 //可以给多个变量赋值，但给出变量后面就要给出值
	var b string = "Dode"
	fmt.Println(a, b, s)
}

// 三：简写赋值 （编译器可以自动决定类型）
func fuzhiJianxie() {
	var a, b, c, d = 1, 2, true, "dode"
	fmt.Println(a, b, c, d)
}

// 四： 短写赋值
func fuzhiDxie() {
	a, b, c, d := 11, 22, true, "dode"
	b = 33 //重新定义变量的值
	fmt.Println(a, b, c, d)
}

// 五： 在函数外定义变量
var (
	aa = 12
	bb = 13
	cc = true
)

func dodeBl() {
	fmt.Println(aa, bb, cc)
}
