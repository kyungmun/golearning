package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Back()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	queue := NewQueue()
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	p := queue.Pop()
	for p != nil {
		fmt.Printf("%v -> ", p)
		p = queue.Pop()
	}

	fmt.Println()
	fmt.Println("ring test")
	r := ring.New(5)
	n := r.Len()

	for i := 0; i < 7; i++ {
		r.Value = 'a' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}

	r.Next().Value = 'k'

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}

	fmt.Println()
	fmt.Println("map test")
	m := make(map[string]string)
	m["1"] = "one"
	m["2"] = "Two"
	m["3"] = "Three"

	for k, v := range m {
		fmt.Printf("map value %s %s\n", k, v)
	}

	delete(m, "1")

	for k, v := range m {
		fmt.Printf("map value %s %s\n", k, v)
	}

}
