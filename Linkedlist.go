package main

import "fmt"

func main() {
	fmt.Println("LinkedList")
	root := node{
		Value: 10,
		Next:  nil,
	}
	head := root.insertion(12)
	// temp := head
	// listprinting(temp)
	// fmt.Println("after insertion at head node")
	head = head.FNI(8)
	// listprinting(head)
	head = head.LNI(14)
	listprinting(head)
	fmt.Println("")
	head = head.FND()
	listprinting(head)
	fmt.Println("")
	head = head.LND()
	listprinting(head)
}

//deletion at mid node

//deletion at last node
func (p *node) LND() *node {
	temp := p
	for temp.Next.Next != nil {
		temp = temp.Next

	}
	temp.Next = nil
	return p
}

//deletion at first node
func (p *node) FND() *node {
	head := p.Next
	p.Next = nil
	return head
}
func (p *node) LNI(data int) *node {
	temp := p
	newnode := createnode(data)
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp = temp.Next
	temp.Next = &newnode
	return p
}

//insertion at forst node of linkedlist
func (p *node) FNI(data int) *node {
	newnode := createnode(data)
	newnode.Next = p
	return &newnode

}

//print the linked list
func listprinting(temp *node) {
	for temp != nil {
		fmt.Print(temp.Value, "->")
		temp = temp.Next
	}
	fmt.Print("nil")
	return

}

//node struct

type node struct {
	Value int
	Next  *node
}

//creation of node
func createnode(data int) node {
	root := node{
		Value: data,
		Next:  nil,
	}
	return root
}

//insertion
func (p *node) insertion(data int) *node {
	newnode := createnode(data)
	p.Next = &newnode
	return p

}
