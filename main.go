package main

import (
	"fmt"

	simplemath "github.com/hanxuema/ddigofunc/simplemath"
)

func main() {
	sv := simplemath.NewVersion(1, 2, 3)
	println(sv.String())
	
	sv.IncrementMajor()
	println(sv.String())

	sv.IncrementMinor()
	println(sv.String())

	sv.IncrementPatch()
	println(sv.String())

	fmt.Printf("%f\n", simplemath.Add(1, 2))

	answer, err := simplemath.Divide(1, 0)
	if err == nil {
		fmt.Printf("%f\n", answer)
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	answer2, _ := simplemath.Divide(1, 0)
	fmt.Printf("%f\n", answer2)

	total := simplemath.Sum(12.12, 12.34, 45.23, 2.2, 44.56, 67.21)
	fmt.Printf("%f\n", total)

	numbers := []float64{12.12, 12.34, 45.23, 2.2, 44.56, 67.21}
	total2 := simplemath.Sum(numbers...)
	fmt.Print("total2", "%f\n", total2)
	fmt.Println("Hello")
}
