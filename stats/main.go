package main

import (
	"fmt"
	"sort"
)

type Statistics struct {
	Mean   float64
	Median float64
	Mode   float64
}

func main() {
	data := []float64{1.0, 6.1, 2, 2, 2, 3, 3.1, 3.4, 5.5, 6.1, 1.1, 2.1}
	s := stats(data)
	fmt.Printf("Mean %v \n Median %v \n Mode %v \n", s.Mean, s.Median, s.Mode)
}

func stats(data []float64) Statistics {
	sort.Float64s(data)
	Mean := mean(data)
	Median := median(data)
	Mode := mode(data)

	return Statistics{
		Mean:   Mean,
		Median: Median,
		Mode:   Mode,
	}
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
	var median float64
	l := len(data)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (data[l/2-1] + data[l/2]) / 2
	} else {
		median = data[l/2]
	}
	return median
}

func mode(data []float64) float64 { // most frequently ***use a map over the frequency and the value // round down pick the on on the left.
	frequency := map[float64]int64{}
	for _, v := range data {
		frequency[v]++
	}
	mode := 0.0
	// TODO: this code is racey. it will return the last detected mode
	// but the order changes because of the map iteration
	//build a tuple type that has value and freq then de dup.
	var highest int64
	for key, value := range frequency {
		if value > highest {
			highest = value
			mode = key
		}
	}

	return mode
}
