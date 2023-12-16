package main

import "fmt"

type Attack interface {
	Attack() string
}

type DB interface {
	GetData() string
	WriteData(data string)
	Closer() error
}
type Hacker struct {
	Name string
	Ip   string
}

func (h *Hacker) Attack() string {
	return fmt.Sprintf("Attack ip is %s", h.Ip)
}

func PrintHack(att Attack) {
	s, ok := att.(*Hacker)
	if ok {
		fmt.Println(s.Name)
		fmt.Println(att.Attack())
	}
}
func main() {
	var hack Hacker
	hack.Ip = "1.1.1.111"
	hack.Name = "kyungmun"
	PrintHack(&hack)
	s := Hacker{"kyungmun2", "21.1.1.1"}
	PrintHack(&s)

}
