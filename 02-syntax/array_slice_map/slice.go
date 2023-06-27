package main

import "fmt"

func main() {
	var array = [...]int{10, 20, 30} // массив
	slice := []int{3, 2, 1}          // слайс
	slice = append(slice, 10)

	fmt.Println(slice[1:4])

	_ = slice[len(slice)-1]
	fmt.Println("array", array, len(array), cap(array))
	fmt.Println("slice", slice, len(slice), cap(slice))

	slice = make([]int, 10, 20)
	fmt.Println("\nslice после инициализации:", slice, len(slice), cap(slice))

	// ассоциативный массив
	var m map[int]string
	//m[3] = "text" // panic: assignment to entry in nil map
	//m = make(map[int]string)
	m = map[int]string{}
	m[3] = "text"
	m[6] = "text"
	for k, v := range m {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}

func sliceCap() {
	slice := make([]int, 0)
	currentCap := cap(slice)
	for i := 0; i < 1_000_000; i++ {
		slice = append(slice, i)
		if cap(slice) != currentCap {
			fmt.Printf("Length: %d. Cap changed from: %d\t, to: %d.\tCooff: %f\n", len(slice), currentCap, cap(slice), float64(cap(slice))/float64(currentCap))
			currentCap = cap(slice)
		}
	}
}
