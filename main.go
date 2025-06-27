package main

import (
	"fmt"
	"time"
)

func main() {
	go ownCount("KSun")
	go ownCount("Peter")
}

func ownCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is counting", i+1)
		time.Sleep(time.Second)
	}
}
