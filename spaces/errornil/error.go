package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "моя ошибка"
}
func validateLength(s string) error {
	if len(s) < 3 {
		return &MyError{}
	}
	return nil
}

func validateLengthBad(s string) error {
	var myErr *MyError
	if len(s) < 3 {
		myErr = &MyError{}
	}
	return myErr
}

func validateLengthFixed(s string) error {
	var myErr *MyError
	if len(s) < 3 {
		myErr = &MyError{}
	}

	if myErr != nil {
		return myErr
	}

	return nil
}

func main() {

	err1 := validateLength("ab")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Ошибки нет")
	}

	err2 := validateLength("abc")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Ошибки нет")
	}

	err3 := validateLengthBad("abc")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("Ошибки нет")
	}
}
