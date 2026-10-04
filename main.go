package main

import (
	"fmt"
	"math"
)

func main() {

	const pow_value = 2
	var height = 1.8
	var weight = 75.0

	var imt = weight / math.Pow(height, pow_value)
	fmt.Println(imt)

	new_value := math.Round(imt)
	fmt.Println(new_value)
}
