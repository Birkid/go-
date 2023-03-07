package main

import "fmt"

// 一：方括号里填写数字，声明有几个变量。
func shuZu() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 11} // 用 ... 让程序自己去数数
	var grid [4][5]int               // 表示4行5列
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}

// 二：range 遍历
func shuZU1() {
	arr := [...]int{2, 3, 4, 9}
	for i := range arr {
		fmt.Println(arr[i])
	}
}

// 三： range 显示行号
func shuZU2() {
	arr := [...]int{2, 3, 4, 9}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 四：range 隐藏行号 用 _ 表示
func shuZU3() {
	arr := [...]int{2, 3, 4, 9}
	for _, v := range arr {
		fmt.Println(v)
	}
}
