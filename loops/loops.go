package main

import (
	"fmt"
)

func main() {

	list := []string{"one", "two", "three"}

	for i := 0; i < 10; i++ {
		fmt.Print("Go!")
	}

	for index, val := range list {
		fmt.Println(index, val)
	}
}
