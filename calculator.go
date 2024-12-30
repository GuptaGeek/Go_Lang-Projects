package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Println("Enter the first name: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second name: ")
	fmt.Scan(&b)
	fmt.Println("Addition:", Addition(a, b))
	fmt.Println("Subtraction:", Subtraction(a, b))
	fmt.Println("Multiplication:", Multiplication(a, b))
	fmt.Println("Divide:", Divide(a, b))
	fmt.Println("Suare:", Square(a))
	fmt.Println("Cube:", Cube(b))

}

type simple interface {
	~int | ~float32 | ~float64
}

// ADDITION

func Addition[T simple](a, b T) T {
	sum := a + b
	return sum

}

// SUBTRACTION
func Subtraction[T simple](a, b T) T {
	return a - b
}

// MULTIPLICATION
func Multiplication[T simple](a, b T) T {
	return a * b
}

// DIVIDE
func Divide(a, b int) int {
	return a / b
}

// SQUARE
func Square(a int) int {
	return a * a
}

// Cube
func Cube(a int) int {
	return a * a * a
}
