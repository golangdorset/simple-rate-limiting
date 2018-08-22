package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// simulate randomish requests by pushing ints into the provided channel
func Feed(f chan int) {
	count := 0
	for {
		randSleep := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(randSleep*10))

		fmt.Println("in:", count)
		f <- count
		count++
	}
}

// simulate randomish requests by pushing ints into the provided channel but pause after 25 request
// to allow the demo "buffer" to refill
func Feed2(f chan int) {
	for i := 0; i < 20; i++ {
		randSleep := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(randSleep*10))

		fmt.Println("in:", i)
		f <- i
	}
	time.Sleep(time.Second * 20)
	for i := 25; i < 100; i++ {
		randSleep := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(randSleep*10))

		fmt.Println("in:", i)
		f <- i
	}
}

// print the time in a simple format
func PrintNow() string {
	return time.Now().Format("15:04:05")
}

// reports the size of the requests queue and burst buffer, and badly pad it to the right of the terminal
func PrintChans(queue chan int, buffer chan bool) {
	for {
		fmt.Printf("\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\trequest queue: %v, buffer: %v\n", len(queue), len(buffer))
		time.Sleep(time.Millisecond * 500)
	}
}
