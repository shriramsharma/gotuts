package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(double(5))
	firstName, lastName := parseName("Shriram Sharma")
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func double(n int) int {
	return n + n
}

func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
