package main

import (
	"errors"
	"fmt"
)

func mySum(x, y int) int {
	return x + y
}

func myPosSub(x, y int) (z int, err error) {
	if x > y {
		z = x - y
	} else {
		err = errors.New("Substraction results in result <= 0!")
	}
	return

}

func myMultiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Printf("mySum(1,2) = %d\n", mySum(1, 2))
	z, _ := myPosSub(20, 5)
	fmt.Printf("myPosSub(20,5) = %d\n", z)
	fmt.Printf("myMultiply(2,2) = %d\n", myMultiply(2, 2))
}
