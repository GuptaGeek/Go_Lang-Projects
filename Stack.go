package main

import "fmt"

func main() {
	fmt.Println("Stack implemenattion")
	array := stack{
		Array: []int{10},
	}
	array.Push(20)
	array.Push(30)
	array.Push(40)
	array.Push(50)
	array.Print()
}

type stack struct {
	Array []int
}

func (p *stack) isEmpty() bool {
	if len(p.Array) == 0 {
		return true
	}
	return false
}
func (p *stack) Peek() {
	i := len(p.Array) - 1
	fmt.Println(p.Array[i])
	return

}

func (p *stack) Push(data int) {
	p.Array = append(p.Array, data)

}
func (p *stack) Pop() {
	if len(p.Array) != 0 {
		i := len(p.Array) - 1
		//fmt.Println(p.Array[i])
		p.Array = p.Array[:i]

	}

}
func (p *stack) Print() {

	for len(p.Array) != 0 {
		p.Peek()
		p.Pop()
	}
}
