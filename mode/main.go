package main

import "fmt"

func main() {
	arr := []int32{6, 1, 4, 4, 4, 5, 3, 0, 0, 0, 4}
	answer := modeFinder(arr)
	fmt.Println(answer)
}

func modeFinder(arr []int32) int32 { //This is A function that take in a  array of int32 and returns the a int32 of the number that accured the most.
	digit := int32(0)
	highest := int32(0)                   //digit and highest are created and set to 0
	uniqueNumber := make(map[int32]int32) // here we create a new empty map with key and value of int32
	for _, v := range arr {               // we range over the array not worrying about the index (_) but settng uniqueNumber's key to the vaue of the arr aray.
		uniqueNumber[v]++
	}
	for k, v := range uniqueNumber { //here we range over the uniqueNumber map setting the key to k and the value to v
		if v > highest { //we check to see if v the value of uniqueNumber is greater the variable highest, that we set to 0 on line 13.
			digit = k //after we range through the uniqueNumber map and find the highest value we know that has accured the most. so we set digit to k and highest to v.
			highest = v
		}
	}
	return digit // then we return the key of our slice digit that is a int32.
}
