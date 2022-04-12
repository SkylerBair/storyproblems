package main

import "fmt"

func main() {
	o := frequency("ass")
	fmt.Println(o)
	fmt.Println(frequency("dylan"))

}

//take a funbction that take a string and returns a map.
func frequency(word string) map[string]int {
	Output := map[string]int{}
	for _, l := range word {
		Output[string(l)]++
	}
	return Output
}
