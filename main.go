package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [4]string{"KSun", "Peter", "John", "Chris"}
	for _, person := range people {
		go isReady(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Received:", <-c)
	}
}

func isReady(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is ready"
}
