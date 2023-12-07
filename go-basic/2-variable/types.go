package main

import "fmt"

func main() {
	var a int = 10
	var b int32 = 20

	fmt.Println(a, b)

	a, b = 1, 2
	var c float32 = 10.34
	fmt.Println(int(c)) //10.34

	fmt.Print(5 % 3) // 2

	var i int32
	var j int64

	i, j = 2, 3

	if i == 2 || j == 3 {
		fmt.Println("equal")
	}

	//boolean
	ok := true
	fmt.Println(ok)

	m := 1

	if m == 1 {
		fmt.Println("is true")
	}

	///////string
	s1 := "hello"
	s2 := "tip js go"

	s := `row1\\r\nrow2`

	fmt.Println(s1, s2, s)

	//concat
	s3 := s1 + s2
	fmt.Println(s3)

	//length
	fmt.Println(len(s3))
	fmt.Println(s3[2:4]) // Cắt chuỗi từ 2 cho tới 4
}
