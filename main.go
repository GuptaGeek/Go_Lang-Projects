package main

import "fmt"

func main() {
	fmt.Println("basic questions")
	fmt.Println(palindromecheck(121))

}
func palindromecheck(n int) bool {
	initialvalue := n
	m := 0
	for {
		r := n % 10
		m = (m * 10) + r
		n = n / 10
		if n == 0 {
			break
		}
	}
	if initialvalue == m {
		return true
	}
	return false
}
