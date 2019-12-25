package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", add(1, 2))

	answer, err := divide(1, 0)
	if err == nil {
		fmt.Printf("%f\n", answer)
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	answer2, _ := divide(1, 0)
	fmt.Printf("%f\n", answer2)

	total := sum(12.12, 12.34, 45.23, 2.2, 44.56, 67.21)
	fmt.Printf("%f\n", total)

	numbers := []float64{12.12, 12.34, 45.23, 2.2, 44.56, 67.21}
	total2 := sum(numbers...)
	fmt.Print("total2", "%f\n", total2)
	fmt.Println("Hello")
}

func sum(values ...float64) float64 {
	result := 0.0
	for _, value := range values {
		result += value
	}
	return result
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func multiply(p1, p2 float64) float64 {
	return p1 * p2
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("can not divided by zero")
	}
	return p1 / p2, nil
}
