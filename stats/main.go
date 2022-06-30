package main

import "fmt"



func main() {
	data := []float64(1.0, 1.0, 1.0, 2,2, 3.4, 5.5, 6.1)
	stats = stats(data)
}

func stats(data []float64) struct {
	mean = mean(data)
	median = median(data)
	mode = mode(data)
}

func mean(data []float64) float64 { // sum / total of its parts
	length := len(data)
	sum := 0.0
	for i := 0; i < length; i++ {
		sum += (data[i])

	}
	avg := (float64(sum)) / (float64(length))
	return avg
}

func median(data []float64) float64 { // is the middle of the sorted parts 

}

func mode(data []float64) float64 { // most frequently ***use a map over the frequency and the value

}