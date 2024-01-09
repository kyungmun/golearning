package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	http.ListenAndServe(":8080", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Content-Type을 스트리밍 데이터로 설정
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")

	// 스트리밍 데이터 생성 및 전송
	var randomData []byte //:= make([]byte, 1024)
	randomData = []byte("test,file,sender \n")
	w.Write(randomData)
	w.(http.Flusher).Flush() // 스트림 플러시
	for i := 0; i < 10; i++ {
		randomData = []byte("test file contents \n")
		w.Write(randomData)
		w.(http.Flusher).Flush() // 스트림 플러시
		time.Sleep(time.Second)  // 1초 대기
	}

}
