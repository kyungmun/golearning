package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance       = 1000
	EarnPoint     = 500
	LosePoint     = 100
	VictoryPoint  = 5000
	GameoverPoint = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//num := rand.Intn(6)
	//fmt.Println(n)
	//cnt := 1
	balance := Balance

	for {
		num := rand.Intn(6)
		fmt.Printf("1 ~ 5 사이의 숫자를 입력하세요 : ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력할 수 있습니다.")

		} else if n < 1 || n > 5 {
			fmt.Println("1 ~ 5 사이의 숫자만 입력하세요.")
		} else {
			if n == num {
				balance += EarnPoint
				fmt.Println("축하합니다. 맞추셨습니다. 남은 돈 : ", balance)
				if balance >= VictoryPoint {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("꽝!! 아쉽지만 다음 기회를...남은 돈 : ", balance)
				if balance <= GameoverPoint {
					fmt.Println("게임 종료")
					break
				}
			}
		}
	}
}
