/*
컴퓨터에 사용할 프린터를 매핑해서 자유롭게 프린터를 변경하면서 사용할수있다.
*/
package main

import (
	"fmt"
)

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Mac Printer")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Windows Printer")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by Epson Printer")
}

type Hp struct{}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by Hp Printer")
}

func main() {
	macComputer := &Mac{}
	winComputer := &Windows{}

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()

}
