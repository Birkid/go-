package main

// 一： 示例一 。创建函数eval，给出定义值和类型，返回一个int类型的值。
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("输入有误!" + op)
	}
}

// 二： 示例二，返回多个值
func div(a, b int) (int, i int) {
	return a / b, a % b
}

// 三： 函数可以有可变数量的参数,用 ... 表示可变参数，用 _ 表示不关系索引
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num // 记不住没关系，total = total + num
	}
	return total
}
