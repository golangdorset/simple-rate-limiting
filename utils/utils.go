package utils

import (
		"math/rand"
	"time"
)

// simulate randomish requests by pushing ints into the provided channel
func Feed(f chan int) {
	count := 0
	for {
		randSleep := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(randSleep*10))

		PrintColOne("in:", count)
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

		PrintColOne("in:", i)
		f <- i
	}
	time.Sleep(time.Second * 20)
	for i := 25; i < 100; i++ {
		randSleep := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(randSleep*10))

		PrintColOne("in:", i)
		f <- i
	}
}
