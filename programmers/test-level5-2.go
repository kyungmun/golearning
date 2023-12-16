package main

/*
https://school.programmers.co.kr/learn/courses/30/lessons/76504?language=go
*/

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("programmer test level5 : city move counter")

	roads := [][]int{{1, 2, 3}, {0, 3, 2}}
	querys := []int64{0, 1, 2, 3, 4, 5, 6}
	z := 5

	result := solution(5, z, roads, querys)

	fmt.Println("result : ", result)
}

//n : 도시의 개수
//z : 초기 도시 그대로 있을시 받는 금액
//roads  : 도시간 도로 정보 {1, 2, 3} 1:출발도시번호, 2: 도착도시번호, 3: 받는 금액
//querys : 만들어야 할 금액 배열
func solution(n int, z int, roads [][]int, queries []int64) []int64 {
	var result = make([]int64, len(queries)-1)
	roadMap := make(map[int]string)

	//도로별 금액 맵
	for _, load := range roads {
		//fmt.Println(load)
		roadMap[load[2]] = fmt.Sprintf("%d-%d", load[0], load[1])
	}
	//fmt.Println(roadMap)

	var wg sync.WaitGroup
	for idx, currency := range queries {
		wg.Add(1)
		go func() {
			var currCity int = 0
			var count int64 = 0
			Calcurator(idx, z, currCity, count, currency, roadMap, result)
			wg.Done()
		}()
		wg.Wait()
	}

	return result
}

func Calcurator(idx, z, currCity int, count, currency int64, roadMap map[int]string, result []int64) {
	//움직이지 않아도 되는 경우
	if currency == int64(z) {
		fmt.Println("same z")
		result[idx] = count //움직이지 않아도 됨
		return
	}

	//남은 금액과 동일한 도로가 있는지 확인
	if w, ok := roadMap[int(currency)]; ok {
		//있으면 현재 도로인지, 다른 도로인지
		startCity := roadMap[int(currency)]
		startCity = strings.Split(startCity, "-")[0] //	출발지 도시 번호
		if num, err := strconv.Atoi(startCity); err != nil {
			if num == currCity { //현재도시와 같으면 1회 이동후 종료
				fmt.Println("same city")
				count += 1
				result[idx] = count
				return
			} else {
				//해당 도시로 이동후 도로 이용 2회 처리후 종료
				count += 2
				fmt.Printf("same %d city\n", num)
				result[idx] = count
				return
			}
		}
	} else { //동일한 도로가 없다면 모아야 하는 금액을 늘려 가면서 금액이 다 차는 경우 찾 아야 함
		//currency = currency - w
		Calcurator(idx, z, currCity, count, currency, roadMap, result)
	}

	count += 1
	if currency == sumMony {
		result[idx] = count
		return
	}
}
