package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("dial")

	buf := make([]byte, 1<<19) //512 KB
	receiveTot := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			} else {
				fmt.Println("read data EOF")
			}
			break
		}
		receiveTot = receiveTot + n
		fmt.Printf("read %d bytes, tot %d \n", n, receiveTot)
	}

	conn.Close()
	fmt.Println("disconnect")
}
