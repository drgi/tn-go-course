package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("s1 allocated:", s1)
	s2 := s1[1:2]
	// нарезка 1 - 2 = длинна элементов
	// s2 на самом деле ссылаеться на массив s1
	fmt.Println("s1 before update s2:", s1)
	s2[0] = 10 // будет изменен слайс s1 !!
	fmt.Println("s1 after update s2:", s1)

	// append создает новый массив, если количество элементов больше cap
	fmt.Println("s1 before append s2:", s1)
	s2 = append(s2, 20) // будет изменен слайс s1 !!
	fmt.Println("s1 after append s2:", s1)

	fmt.Println("s1 before append s2:", s1)
	s2 = append(s2, 20, 30, 40, 50, 60, 70) // НЕ будет изменен слайс s1 !!
	fmt.Println("s2 after append:", s2)
	fmt.Println("s1 after append s2:", s1)
}
