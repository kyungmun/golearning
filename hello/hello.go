package main

import (
	"fmt"
	"sort"

	"goproject/hello/argument"
)

type User struct {
	Name    string
	Address string
}
type Product struct {
	User
	Name         string
	Price        int
	ReviwerScore float64
}

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

type Monster struct {
	Actor
	Attack int
	Speed  int
}

type Player struct {
	Name string
	Age  int
	Goal int
	Pass float64
}

type Players []Player

func (s Players) Len() int {
	return len(s)
}

func (s Players) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Players) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

const (
	Red int = iota
	Blue
	Green
	Black
)

func Divide(a, b int) (int, bool) {
	const aa = 10
	return aa, true
}

var slice = make([]int, 3, 5)
var array = [10]int{1, 1, 1}

func main() {
	var test = "string"
	fmt.Println(test)
	var monster = Monster{Actor{"lim", 100, 8.7}, 500, 200}
	fmt.Println(monster.Speed)
	fmt.Println(monster.Actor.Speed)

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	} else {
		fmt.Println("slice length", len(slice))
		for i := 0; i < len(slice); i++ {
			//fmt.Println(slice[i])
			slice[i] = i + 1
		}
		slice2 := append(slice, 60, 70)
		for i, v := range slice {
			fmt.Printf("%d\t %d\n", i, v)
		}
		slice[1] = 100
		for i, v := range slice2 {
			fmt.Printf("%d\t %d\n", i, v)
		}
		slice2 = append(slice2[:2], slice2[2+1:]...)
		slice4 := make([]int, 5)

		copy(slice4, slice2)
		sort.Ints(slice2)

		fmt.Println(slice2)
		for i, v := range slice2 {
			fmt.Printf("%d\t %d\n", i, v)
		}

		s := []Player{
			{"나통키", 13, 45, 78.4},
			{"오명태", 16, 24, 67.4},
			{"오동도", 18, 54, 58.8},
			{"황금산", 17, 36, 89.7},
		}
		//sort.Sort(sort.Reverse(Players(s)))
		sort.Sort(Players(s))
		fmt.Println(s)
	}

	defer fmt.Println(argument.Sum(1, 2, 3, 4, 5))
	defer fmt.Println(argument.Sum(1, 2, 5))

}
