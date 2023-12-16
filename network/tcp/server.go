package main

import (
	"crypto/rand"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept : ", err)
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

	payload := make([]byte, 1<<24) //16MB
	fmt.Printf("write data buffer length %d \n", len(payload))
	_, err := rand.Read(payload)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("connect %v \n", conn.RemoteAddr())
	defer conn.Close()

	_, err = conn.Write(payload)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("write done")

}
