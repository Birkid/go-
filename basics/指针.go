package main

// 一： 改变变量的值，要传指针进去。
func swap(a, b *int) {
	*a, *b = *b, *a
	// 利用 & 接回地址
	/*	a, b := 3, 4
		swap(&a, &b)*/
}

// 二： 利用return 返回
func swap1(a, b int) (int, i int) {
	return b, a
	/*	a, b := 3, 4
		a, b = swap1(b, a)*/
}
