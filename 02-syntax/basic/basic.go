package basic

import "fmt"

func Range() {
	slice := []int{10, 20, 30, 40}
	arr := [...]int{10, 20, 30, 40}
	str := "string"

	for i, v := range slice {
		fmt.Println("Index: ", i, "Value: ", v)
	}

	for i, v := range arr {
		fmt.Println("Index: ", i, "Value: ", v)
	}

	for i, v := range str {
		fmt.Println("Index: ", i, "Value: ", v)
	}
}

func Vars() {
	var i int
	var u uint = 10
	x := 20
	fmt.Println("i: ", i, "u: ", u, "x: ", x)
}

func Pointers() {
	var s string = "ABS"
	var pointer *string = &s
	fmt.Println(pointer, &pointer)
}

func Scopes() {
	x := 1
	{
		x := 2
		fmt.Println("In scope x: ", x)
	}
	fmt.Println("x: ", x)
}
