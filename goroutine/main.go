package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SystemResource struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
}

func main() {
	r := gin.Default()

	r.GET("/system/resources", func(c *gin.Context) {
		handleTCPConnection(c)
	})

	r.Run(":8080") // 서버를 8080 포트에서 실행
}

func handleTCPConnection(c *gin.Context) {
	// TCP 연결을 생성하여 b 프로세스에 연결
	// conn, err := net.Dial("tcp", "localhost:8888") // 예시로 localhost:8888로 가정
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to backend"})
	// 	return
	// }
	// defer conn.Close()

	// 데이터 수신 및 쌓기 처리 고루틴 실행
	receiveAndProcessData(nil, c)

	fmt.Println("response finish")
	c.Status(http.StatusOK)

}

func receiveAndProcessData(conn net.Conn, c *gin.Context) {
	// 데이터를 쌓을 큐 채널 생성
	dataQueue := make(chan []byte, 100) // 최대 100개의 데이터를 보관할 수 있는 큐

	var wg sync.WaitGroup

	wg.Add(1)
	// 데이터를 수신하고 큐에 저장
	go func() {
		defer wg.Done()
		idx := 0
		//buffer := make([]byte, 1024)
		for {
			// _, err := conn.Read(buffer)
			// if err != nil {
			// 	fmt.Println("Error reading data:", err)
			// 	close(dataQueue)
			// 	return
			// }
			// 큐에 데이터 저장
			idx++
			buffer := []byte(fmt.Sprintf(`{"cpu_usage":%d.%d, "memory_usage":%d.%d}`, idx, idx, idx, idx))
			// 큐에 데이터 저장
			dataQueue <- buffer
			if idx == 10000 {
				fmt.Println("receive data finish")
				dataQueue <- []byte("exit")
				break
			}
		}
	}()

	fmt.Printf("goroutine start \n")
	// 요청 컨텍스트에 대한 뮤텍스
	var ctxMutex sync.Mutex
	// 데이터를 쌓아서 처리하고 전송하는 고루틴 실행
	for i := 0; i < 3; i++ { // 예시로 5개의 고루틴으로 실행
		go processDataAndSend(i, dataQueue, c, &ctxMutex, &wg)
	}
	wg.Wait()
}

func processDataAndSend(idx int, dataQueue chan []byte, c *gin.Context, ctxMutex *sync.Mutex, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for data := range dataQueue {
		//fmt.Printf("goroutine idx : %d, %v\n", idx, string(data))
		// 특정 조건을 확인하여 루프를 종료합니다.
		if bytes.Equal(data, []byte("exit")) {
			fmt.Println("for end")
			close(dataQueue)
			break
		}
		// 클라이언트에 데이터 전송
		err := sendResponse(c, &data, idx, ctxMutex)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
	fmt.Printf("send goroutine %d end\n", idx)
}

// 클라이언트에 데이터를 전송하는 함수
func sendResponse(c *gin.Context, resource *[]byte, idx int, ctxMutex *sync.Mutex) error {
	ctxMutex.Lock()
	defer ctxMutex.Unlock()
	//fmt.Printf("goroutine %d write, %v\n", idx, string(*resource))

	// 클라이언트에 JSON 형식의 데이터 전송
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// 데이터를 JSON 형식으로 변환
	// jsonData, err := json.Marshal(resource)
	// if err != nil {
	// 	return err
	// }
	c.Writer.Write(*resource)

	return nil
}
