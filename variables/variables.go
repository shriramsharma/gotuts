package main

import (
	"fmt"
	"io/ioutil"
	m "math"
	"net/http"
	"strconv"
)

func main() {
	var one int = 1
	fmt.Println(one)

	var two, three string = "two", "three"
	fmt.Println(two, three)

	var four = true
	fmt.Println(four)

	five := "I know what I am"
	fmt.Println(five)

	sum, prod := learnMultiple(3, 4)
	fmt.Println("Sum: ", sum)
	fmt.Println("Prod: ", prod)

	bs := []byte("a slice")
	fmt.Println(bs)

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	fmt.Println(m)
	learnFlowControl()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnMemory() (p, q *int) {
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func learnFlowControl() {
	x := 42.0

	switch x {
	case 0:
	case 1:
	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d", key, value)
	}

	if y := expensiveComputation(); y > x {
		x = y
	}

	fmt.Println("Value of x:", x)

	xBig := func() bool {
		return x > 10000
	}

	fmt.Println("xBig: ", xBig())

	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	goto love

love:
	learnFunctionFactory()
	learnDefer()
	learnInterfaces()
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFunctionFactory() {
	fmt.Println(sentenceFactory("summer")("A beautiful", "day"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))

}

func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after)
	}
}

func learnDefer() bool {

	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")

	return true
}

type Stringer interface {
	StringG() string
}

type pair struct {
	x, y int
}

func (p pair) StringG() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{3, 4}
	fmt.Println(p.StringG())
	var i Stringer
	i = p
	fmt.Println(i.StringG())
	fmt.Println(p)
	fmt.Println(i)

	learnVariadicParams("great", "learning", "here")
}

func learnVariadicParams(myStrings ...interface{}) {
	for _, param := range myStrings {
		fmt.Println("param: ", param)
	}

	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("no one there")
	} else {
		fmt.Println(x)
	}

	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}

	learnConcurrency()
}

func inc(i int, c chan int) {
	c <- i + 1
}

func learnConcurrency() {
	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c)

	cs := make(chan string)
	ccs := make(chan chan string)
	go func() { c <- 84 }()
	go func() { cs <- "wordy" }()

	select {
	case i := <-c:
		fmt.Println("it's a %T", i)
	case <-cs:
		fmt.Println("its a string")
	case <-ccs:
		fmt.Println("didnt happen")
	}

	learnWebProgramming()

}

func learnWebProgramming() {
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err)
	}()

	requestServer()
}

func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You learned Go in Y minutes"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: '%s'", string(body))
}
