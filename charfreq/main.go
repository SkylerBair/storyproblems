package main

import "fmt"

func main() {
	frequency("mississippi")
	fmt.Println(Output)
}

var Output = map[string]int{}

//take a funbction that take a string and returns a map.
func frequency(word string) map[string]int {
	for _, l := range word {
		Output[string(l)]++
	}
	return Output
}
