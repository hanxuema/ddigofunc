package main

import "fmt"

func main()  {
	fmt.Printf("%f\n", add(1,2))
	fmt.Println("Hello")
}

func add(p1, p2 float64) float64 {
	
	return p1 + p2
}