package main

/*
https://school.programmers.co.kr/learn/courses/30/lessons/49190
*/

import (
	"fmt"
)

func move(x, y, arrow int) (int, int, string) {
	switch arrow {
	case 0:
		return x, y + 1, inttostr(x, y) + inttostr(x, y+1)
	case 1:
		return x + 1, y + 1, inttostr(x, y) + inttostr(x+1, y+1)
	case 2:
		return x + 1, y, inttostr(x, y) + inttostr(x+1, y)
	case 3:
		return x + 1, y - 1, inttostr(x, y) + inttostr(x+1, y-1)
	case 4:
		return x, y - 1, inttostr(x, y-1) + inttostr(x, y)
	case 5:
		return x - 1, y - 1, inttostr(x-1, y-1) + inttostr(x, y)
	case 6:
		return x - 1, y, inttostr(x-1, y) + inttostr(x, y)
	case 7:
		return x - 1, y + 1, inttostr(x-1, y+1) + inttostr(x, y)
	default:
		return x, y, inttostr(x, y) + inttostr(x, y)
	}
}

func inttostr(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func solution(arrows []int) int {
	var movepoint = make(map[string]int)
	var checkpoint = make(map[string]int)
	var x, y, result int = 0, 0, 0
	var mx, my, oldx, oldy int = 0, 0, 0, 0
	var node, edge, cross_edge string
	movepoint[inttostr(x, y)] = 1
	for _, v := range arrows {
		oldx = x
		oldy = y
		mx, my, edge = move(x, y, v)
		cross_edge = inttostr(0, 0)
		switch v {
		case 1:
			cross_edge = inttostr(oldx, oldy+1) + inttostr(mx, my-1)
		case 3:
			cross_edge = inttostr(oldx, oldy-1) + inttostr(mx, my+1)
		case 5:
			cross_edge = inttostr(mx, my+1) + inttostr(oldx, oldy-1)
		case 7:
			cross_edge = inttostr(mx, my-1) + inttostr(oldx, oldy+1)
		}

		x = mx
		y = my

		node = inttostr(mx, my)

		if _, ok := checkpoint[edge]; !ok {
			if _, ok := checkpoint[cross_edge]; ok {
				result += 1
				fmt.Printf("cross_edge value : %s\n", cross_edge)
			}
			if _, ok := movepoint[node]; ok {
				result += 1
				fmt.Printf("node value : %s\n", node)
			} else {
				movepoint[node] = 1
			}
			checkpoint[edge] = 1
		}
	}
	//for k, _ := range checkpoint {
	//	fmt.Printf("value : %s\n", k)
	//}
	//fmt.Println(movepoint)
	//fmt.Println(checkpoint)

	return result
}

func main() {
	fmt.Println("programmer test level5 : room counter")

	//value := []int{6, 6, 6, 4, 4, 4, 2, 2, 2, 0, 0, 0, 1, 6, 5, 5, 3, 6, 0}
	//value := []int{6, 4, 2, 0, 0, 6, 6, 4, 4, 4, 2, 2, 2, 0, 0, 0, 6, 4, 2}
	value := []int{5, 2, 2, 7, 7, 2, 2, 5, 1, 5}
	result := solution(value)
	fmt.Println(result)

}
