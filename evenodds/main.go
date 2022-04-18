package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens, odds := oddsAndEvens(s)
	fmt.Println("Evens: ", evens, "\nOdds: ", odds)

}

func oddsAndEvens(s []int) ([]int, []int) {
	evens := []int{}
	odds := []int{}
	for _, s := range s {
		s1 := s % 2
		if s1 != 0 {
			evens = append(evens, s)

		} else {
			odds = append(odds, s)

		}
	}
	return odds, evens
}
