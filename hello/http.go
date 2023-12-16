package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func chanWriter(opChan chan<- string) {
	for index := 0; index < 100; index++ {
		opChan <- fmt.Sprintf("%d = 123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", index)
		time.Sleep(1 * time.Second)
	}
	//close(opChan)
}

func ginServer(addr string) {
	api := gin.Default()
	api.GET("/:id", func(c *gin.Context) {
		//w := bufio.NewWriterSize(c.Writer, 10)

		taskID := c.Param("id")
		fmt.Println(taskID)
		opChan := make(chan string)
		//defer close(opChan)
		go chanWriter(opChan)

		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if msg, ok := <-opChan; ok {
				//c.Writer.Header().Set("Content-Type", "text/event-stream")
				fmt.Println(msg)
				c.Writer.Write([]byte(msg + "\n"))
				//c.SSEvent("message", msg)
				return true
			}
			return false
		})
		/*
			//go func() {
			for {
				data1 := <-opChan
				// c.Writer.Write([]byte(taskID + data1))
				// c.Writer.Write([]byte("/0x00"))
				//c.String(http.StatusOK, taskID+data1)
				//c.Writer.Write([]byte("0x10"))
				//c.Writer.Write([]byte(data1))
				//w.Write([]byte("0x00"))
				//c.Writer.Flush()
				c.SSEvent("data", data1)
				// if f, ok := c.Writer.(http.Flusher); ok {
				// 	f.Flush()
				// }
				// case <-time.After(5 * time.Second):
				// 	fmt.Println("timeout")
				// 	c.String(http.StatusRequestTimeout, "Timeout")
				// 	close(opChan)
				// 	return
				// }
			}
			//}()
		*/
		//c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		//c.Writer.WriteHeader(http.StatusOK)

		//opChan <- "test1"
		//opChan <- "test2"
		/*
			c.Stream(func(w io.Writer) bool {
				output, ok := <-opChan
				if !ok {
					return false
				}
				outputBytes := bytes.NewBufferString(taskID + output)
				c.Writer.Write(append(outputBytes.Bytes(), []byte("\n")...))
				return true
			})
		*/
	})

	api.Run(addr)
}

func httpServer(addr string) {
	srv := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		opChan := make(chan string)
		go chanWriter(opChan)
		w.WriteHeader(200)
		for output := range opChan {
			outputBytes := bytes.NewBufferString(output)
			w.Write(append(outputBytes.Bytes(), []byte("\n")...))
			w.(http.Flusher).Flush()

		}
	})

	srv.ListenAndServe()

}

func main() {

	go ginServer(":8081")
	/*
		go httpServer(":8082", chCmd)

		for {
			select {
			case msg1 := <-chCmd:
				fmt.Println(msg1)
			}
		}
	*/
	<-make(chan interface{})
}
