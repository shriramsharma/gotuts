package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {
	i := 20
	fmt.Println("Value of i is", i)
	zeroval(i)
	fmt.Println("Value of i after calling zeroval is ", i)
	zeroptr(&i)
	fmt.Println("Value of i after calling zeroptr is ", i)
}
