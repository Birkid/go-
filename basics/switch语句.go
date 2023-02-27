package main

// 一： 示例一
/*
func main() {
	num := 3 			// 声明一个变量 num 等于3
	switch num {
	case 1:			 	// cash关键字，匹配到哪个就执行哪个。
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}

*/
/*
// 二： 示例二
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score : %d", score))
	case score < 60:
		g = "a"
	case score < 70:
		g = "b"
	case score < 90:
		g = "c"
		//default:
		//	fmt.Println("wrong")

	}
	return g
}
func main() {
	fmt.Println(
		grade(60),
		grade(30),
		grade(88),
		grade(101),
	)

}

*/
/*
// 三： 示例三 支持没有表达式的形式，直接对多个条件进行匹配。
func main() {
	age := 20

	switch {
	case age < 18:
		fmt.Println("You are under 18 years old.")
	case age >= 18 && age <= 60:
		fmt.Println("You are between 18 - 60 ")

	}
}

*/
/*
// 四： 示例四 使用多个cash值
func main() {
	fruit := "orange"

	switch fruit {
	case "apple", "orange", "banana":
		fmt.Println("This is a fruit.")
	default:
		fmt.Println("Unknown fruit.")
	}
}

*/
