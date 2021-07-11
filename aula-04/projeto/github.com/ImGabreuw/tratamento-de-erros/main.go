package main

import (
	"errors"
	"fmt"
)

func main() {

	result, _ := soma(7, 2)

	fmt.Println(result)

}

func soma(x int, y int) (int, error) {
	result := x + y

	if result > 10 {
		return 0, errors.New("total maior que 10")
	}

	return result, nil
}