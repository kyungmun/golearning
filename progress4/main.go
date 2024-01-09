package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	escape           = "\033"
	reset            = escape + "[0m"
	red              = escape + "[31m"
	green            = escape + "[32m"
	yellow           = escape + "[33m"
	skyblue          = escape + "[34m"
	progressBarWidth = 50
)

func main() {
	// ... (이전 코드와 동일)

	// TODO: 서버로부터 데이터 조회 및 파일 저장

	// TODO: 프로그래스바 처리
	fmt.Println("작업 진행 중...")
	progressChannel := make(chan int)
	defer close(progressChannel)
	total := 100 // 예시로 100개의 작업을 가정합니다.

	go simulateDataRetrieval(total, progressChannel)
	showProgressBar(total, progressChannel)
	//go simulateDataRetrieval(total, progressChannel)
	//showProgressBar(total, progressChannel)

	fmt.Println("작업 완료")
}

func showProgressBar(total int, progressChannel chan int) {
	terminalWidth, _, err := getTerminalWidth()
	if err != nil {
		fmt.Println("터미널 넓이를 가져오는 중 오류 발생:", err)
		return
	}

	//var prevProgressBar string

	for i := 0; i < total; i++ {
		newWidth, newHeight, newErr := getTerminalWidth()
		if newErr == nil && newWidth != terminalWidth {
			terminalWidth, _ = newWidth, newHeight
		}
		progress := <-progressChannel
		percentage := int(float64(progress) / float64(total) * 100)
		barLength := int(float64(terminalWidth) * 0.8) // 프로그래스바를 터미널 창의 80%로 제한

		color := green
		if percentage < 50 {
			color = yellow
		} else if percentage < 80 {
			color = skyblue
		}

		equalsCount := int(float64(barLength) * float64(percentage) / 100)
		spacesCount := barLength - equalsCount
		equals := fmt.Sprintf("%s=%s", color, reset)
		progressBar := strings.Repeat(equals, equalsCount) + ">" + strings.Repeat("-", spacesCount)

		fmt.Printf("\rDownloading [%s%s] %d%% %s", progressBar, reset, percentage, reset)

		// 현재 프로그래스바를 출력할 위치를 계산
		// cursorPosition := fmt.Sprintf("\033[%d;%dH", terminalHeight, 0)

		// // 이전에 출력된 프로그래스바를 지우는 부분
		// clearPrevProgressBar := fmt.Sprintf("\033[%dA", strings.Count(prevProgressBar, "\n"))

		// // 프로그래스바 출력
		// fmt.Print(cursorPosition + clearPrevProgressBar + "[" + progressBar + "]" + fmt.Sprintf(" %d%%", percentage))

		// // 현재 프로그래스바를 저장
		// prevProgressBar = "[" + progressBar + "]" + fmt.Sprintf(" %d%%", percentage)

		// 화면을 갱신할 시간 지연 (테스트 및 조절 가능)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

func getTerminalWidth() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var width, height int
	fmt.Sscanf(string(out), "%d %d", &height, &width)
	return width, height, nil
}

func simulateDataRetrieval(total int, progressChannel chan int) {
	for i := 0; i < total; i++ {
		// Simulate data retrieval delay
		time.Sleep(time.Millisecond * 100)
		progressChannel <- i + 1
	}
	//close(progressChannel)
}
