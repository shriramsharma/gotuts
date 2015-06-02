package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("Working....")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	bugMsgs := make(chan string, 2)

	bugMsgs <- "buffered"
	bugMsgs <- "msg"

	fmt.Println(<-bugMsgs)
	fmt.Println(<-bugMsgs)

	done := make(chan bool)
	go worker(done)
	<-done

	fmt.Println("I am just testing this")

}
