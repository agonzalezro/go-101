package main

import (
	"fmt"
	"time"
)

type Message struct {
	Content string
}

func produce(msgs chan<- Message, done chan<- bool) {
	for i := 1; i <= 10; i++ {
		msgs <- Message{fmt.Sprintf("Message #%d", i)}
	}
	done <- true
}

func consume(msgs <-chan Message) {
	for {
		select {
		case msg := <-msgs:
			fmt.Println(msg.Content)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	msgs := make(chan Message)
	done := make(chan bool)

	go produce(msgs, done)
	go consume(msgs)

	<-done
}
