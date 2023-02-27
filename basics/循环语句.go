package main

import "fmt"

// 一： 初始化语句在循环开始前执行，条件语句在每次迭代开始前求值，后置语句在每次迭代末尾执行。循环会一直执行，直到条件语句的结果为 false。
func xunHuan1() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

// 二：利用for循环实现while循环的功能 （go中没有while关键字）。
func xunHuan2() {
	sum := 0
	i := 1
	for i <= 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

// 三： 死循环
func xunHuan3() {
	for {
		fmt.Print("Enter command: ")
		var command string
		fmt.Scanln(&command) //获取用户输入并保存到上面定义的字符串变量command中
		if command == "quit" {
			break
		}
		fmt.Println("You entered:", command)
	}
	fmt.Println("Goodbye!")
}
