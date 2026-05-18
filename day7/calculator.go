package main

import (
	"errors"
)

func Add(a, b float64) float64 {
	return a + b

}

func Substract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("На ноль делить нельзя")
	}

	return a / b, nil
}
