package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"KSun", "Peter"}
	for _, person := range people {
		go isReady(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isReady(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person, "is ready")
	c <- true
}
