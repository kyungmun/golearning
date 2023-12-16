package main

/*
프로그래머스 문제
https://school.programmers.co.kr/learn/courses/30/lessons/92334
*/

import (
	"fmt"
	"strings"
)

//내가 푼 방식
func solution(id_list []string, report []string, k int) []int {
	reportMap := make(map[string]int)
	reportMap2 := make(map[string][]string)
	destMap := make(map[string]int)
	stopdestMap := make(map[string]int)
	var returnArr []int

	for _, r := range report {
		use_id := strings.Split(r, " ")[0]
		dest_id := strings.Split(r, " ")[1]
		if use_id == dest_id {
			continue
		}
		fmt.Println(use_id, dest_id)
		//신고자-대상자 값 중복 제거 하면서, k회이상 신고당한 id 기록
		if _, ok := reportMap[r]; !ok {
			fmt.Printf("중복 없는 리포트 추가 : %s\n", r)
			reportMap[r] = 1
			reportMap2[dest_id] = append(reportMap2[dest_id], use_id)
			//신고당한 id 누적건수
			destMap[dest_id] = destMap[dest_id] + 1
			//신고당한 id 의 누적건수가 k 이상(계정정지) 대상인 id 만
			if destMap[dest_id] >= k {
				stopdestMap[dest_id] = destMap[dest_id]
			}
		}
	}
	fmt.Printf("중복없는 리포트 최종 : %d, %v\n", len(reportMap), reportMap)
	fmt.Printf("신고당한 id를 신고한 id : %d, %v\n", len(reportMap2), reportMap2)
	fmt.Printf("stopdestMap count : %d, %v\n", len(stopdestMap), stopdestMap)

	resultMap := make(map[string]int)
	for k1, _ := range stopdestMap {
		for _, id := range reportMap2[k1] {
			resultMap[id] = resultMap[id] + 1
		}
	}
	fmt.Printf("resultMap count : %d, %v\n", len(resultMap), resultMap)

	for _, v := range id_list {
		fmt.Printf("idlist : %s %d\n", v, resultMap[v])
		returnArr = append(returnArr, resultMap[v])
	}

	return returnArr
}

//다른 사람이 푼 방식
func solution2(id_list []string, report []string, k int) []int {
	a := map[string]map[string]bool{}
	var r = make([]int, len(id_list))

	for _, v := range report {
		c := strings.Split(v, " ")
		if _, ok := a[c[1]]; !ok {
			a[c[1]] = make(map[string]bool)
		}
		a[c[1]][c[0]] = true
	}
	fmt.Printf("%v", a)

	for _, v := range a {
		if len(v) > k-1 {
			for k, _ := range v {
				for i, n := range id_list {
					if n == k {
						r[i]++
					}
				}
			}
		}
	}

	return r
}

func main() {
	//id_list := []string{"kyungmun", "mun", "lim"}
	//report := []string{"kyungmun lim", "lim mun", "kyungmun mun", "kyungmun lim", "kyungmun mun"}
	//id_list := []string{"con", "ryan"}
	//report := []string{"ryan con", "ryan con", "ryan con", "ryan con"}
	id_list := []string{"muzi", "frodo", "apeach", "neo"}
	report := []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}
	result := solution(id_list, report, 2)

	fmt.Println(result)
}
