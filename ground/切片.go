package main

import "fmt"

func sliceDemo() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
	s := arr[1:6]
	s[0] = 10
	fmt.Println(s)
	s1 := arr[:]
	fmt.Println(s1)
	s2 := s1[:5]
	s3 := s2[2:]
	fmt.Println(s3)

}
