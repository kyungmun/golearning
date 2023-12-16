package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntToBuffer(t *testing.T) {
	payload := make([]byte, 1<<24) //16MB
	t.Logf("write data buffer length %d", len(payload))
	_, err := rand.Read(payload)
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			t.Fatal(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1<<19) //512 KB
	receiveTot := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			} else {
				t.Log("read data EOF")
			}
			break
		}
		receiveTot = receiveTot + n
		t.Logf("read %d bytes, tot %d", n, receiveTot)
	}

	conn.Close()

}
