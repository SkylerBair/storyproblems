package main

import "fmt"

func main() {
	fib := fib_sequence(0)
	fmt.Println("0 fibonacci sequence's =", fib)

	fib1 := fib_sequence(1)
	fmt.Println("1 fibonacci sequence's =", fib1)

	fib2 := fib_sequence(2)
	fmt.Println("2 fibonacci sequence's =", fib2)

	fib3 := fib_sequence(3)
	fmt.Println("3 fibonacci sequence's =", fib3)

	fib4 := fib_sequence(4)
	fmt.Println("4 fibonacci sequence's =", fib4)

	fib5 := fib_sequence(5)
	fmt.Println("5 fibonacci sequence's =", fib5)

	fib6 := fib_sequence(6)
	fmt.Println("6 fibonacci sequence's =", fib6)
}

func fib_sequence(n int) []int {
	one := []int{0, 1, 1}
	numbers := []int{0, 1}
	for i := 1; i <= n; i++ {
		if n == 0 {
			break
		} else if n == 1 {
			return one
		} else if n > 1 {
			first := numbers[len(numbers)-2]
			last := numbers[len(numbers)-1]
			ans := first + last
			numbers = append(numbers, ans)
		}
	}
	return numbers
}
