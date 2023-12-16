package main

import "fmt"

type account struct {
	blance int
}

type DuckInterfase interface {
	Fly()
	Walk(distance int) int
}

type Sample interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("hi, im %d %s", s.Age, s.Name)
}

func withdrawFunc(a *account, amount int) {
	a.blance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.blance -= amount
}
func (a *account) withdrawMethod2(amount int) {
	a.blance += amount
}
func (a account) withdrawMethod3(amount int) account {
	a.blance += amount
	return a
}

func main() {
	student := Student{"kyungmun", 48}
	var stringer Sample
	stringer = student

	fmt.Printf("%s\n", stringer.String())
	a := &account{100}
	withdrawFunc(a, 30)
	fmt.Printf("%d \n", a.blance)

	a.withdrawMethod(30)
	a.withdrawMethod(100)
	a.withdrawMethod2(50)
	//*a = a.withdrawMethod3(10)
	fmt.Printf("%d \n", a.blance)
}
