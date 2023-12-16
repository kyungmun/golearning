package main

import "fmt"

func changeArray(array2 *[5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	fmt.Println("slice sample")
	var array = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}

	changeArray(&array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)

	addNum(slice)
	fmt.Println(slice)

	//슬라이싱 해도 복사되는게 아닌 슬라이싱하는 원래값의 데이터 포인터를 갖는다.
	slice2 := slice[1:4]
	fmt.Println(slice2)

	slice[2] = 77

	fmt.Println(slice)
	fmt.Println(slice2)

	slice3 := append([]int{}, slice...)
	fmt.Println(slice3)
	slice4 := make([]int, len(slice))
	copy(slice4, slice)
	fmt.Println(slice4)
}

func addNum(slice []int) {
	slice = append(slice, 4)
}
