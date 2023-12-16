package main

import (
	"fmt"
	"math"
)

type Data struct {
	A int
	B int
}

func (d *Data) Calculator() int {

	value := math.Pow(float64(d.A), float64(d.B))
	fmt.Printf("%f\n", value)
	return int(uint64(value) / 100)
}

func main2() {
	var count, a, b int
	fmt.Print("테스트 케이스 개수를 입력하십시오. ")
	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println(err)
	} else {
		dataArr := make([]Data, count)
		for i := 0; i < count; i++ {
			fmt.Printf("%d 번째 테스트에 사용될 값 두 수를 입력하십시오. ", i+1)
			fmt.Scanln(&a, &b)
			dataArr[i].A = a
			dataArr[i].B = b
			fmt.Printf("%d %d\n", dataArr[i].A, dataArr[i].B)
		}

		fmt.Println("---result---")
		for _, data := range dataArr {
			fmt.Printf("%d ** %d %% 10 = %d \n", data.A, data.B, data.Calculator())
		}
	}

}
