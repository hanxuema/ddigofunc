package simplemath

import (
	"errors"
	"math"
)

func Sum(values ...float64) float64 {
	result := 0.0
	for _, value := range values {
		result += value
	}
	return result
}

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("can not divided by zero")
	}
	return p1 / p2, nil
}

func privateFun() {

}
