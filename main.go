package main

import (
	"log"
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func holaMundo(t time.Time) {
	log.Println("Hello Go World!")
}

func main() {
	doEvery(5*time.Second, holaMundo)
}
