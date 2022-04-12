package main

import "fmt"

const oneDollar uint8 = 1.00

type dollars struct {
	oneDollar        uint8
	fiveDollar       uint8
	tenDollar        uint8
	twentyDollar     uint8
	fiftyDollar      uint8
	oneHundredDollar uint16
}

type cents struct {
	penny  uint64
	nickle uint8
	dime   uint8
	quater uint8
}

func main() {

	total := 400.00
	providedAmount := 10.00

	change := changeCalculator(total, providedAmount)
	fmt.Println(change)

	cashCheck(total)

}

// func changeConverter(uint8) string {
// 	c =
// }

func changeCalculator(total float64, providedAmount float64) float64 {
	changeInActual := 0.00
	changeInActual = providedAmount - total
	return changeInActual
}

func cashCheck(total float64) {
	if total >= 1 {
		t := total / 1
		fmt.Printf("%v $Dollar Bills \n", t)
		t = t * 1

		fmt.Println(t)
	}

}
