package main

import (
	"github.com/golangdorset/simpleratelimiting/utils"
	"time"
)

func main() {
	one()

	//two()

	//three()
}

// ranging over the channel you can see that we print out the int as soon as it is written in by the feeder.
func one() {
	requests := make(chan int, 100)

	go utils.Feed(requests)

	for i := range requests {
		utils.PrintOut("out:", i)
	}
}

// using a ticker that puts an object on the channel every 1 second, we can delay the reading from the requests queue
// as shown by the rate of requests going in, and being read out
func two() {
	requests := make(chan int, 100)

	go utils.Feed(requests)

	drip := time.Tick(time.Second)

	for i := range requests {
		<-drip
		utils.PrintOut("out:", i, utils.PrintNow())
	}
}

// feed2 randomly adds up to 20 ints at random intervals, then will sleep for 20 seconds before continuing, this
// allows the "buffer" to re-fill.
// to start with the requests are processed at the same speed they are received, but once the burst of 5 has been used,
// we are back to 1 a second.
func three() {
	requests := make(chan int, 100)

	go utils.Feed2(requests)

	drip := time.Tick(time.Second)

	//type of this channel is not important
	burstyLimiter := make(chan bool, 5)

	// pre-fill the buffer channel so that the first 5 requests can be processes without waiting for the "tick"
	for i := 0; i < 5; i++ {
		burstyLimiter <- true
	}

	// in the background, every second (using the drip feed ticker from before) add a "token" to the bursty limiter
	// if no requests come in for a period of time, then this "buffer" will fill back to 5
	// the next 5 requests will be able to read from the channel and be processed straight away
	go func() {
		for {
			<-drip
			burstyLimiter <- true
		}
	}()

	// print the status of the request queue and size of the "buffer" limiter
	go utils.PrintChans(requests, burstyLimiter)

	for i := range requests {
		<-burstyLimiter
		utils.PrintOut("out:", i, utils.PrintNow())
	}
}
