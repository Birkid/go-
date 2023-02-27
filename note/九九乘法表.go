package main

import "fmt"

func chengFa() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

}

// 外层循环控制行数，内层循环控制列数。每行的列数逐渐递增，第一行只有一列，第二行有两列，以此类推。
