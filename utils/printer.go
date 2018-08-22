package utils

import (
	"fmt"
	"time"
)

const format = "%-25s  %-25s  %-25s\n"

func Print(in, out, status string) {
	fmt.Printf(format, in, out, status)
}

func PrintColOne(s ...interface{}) {
	str := fmt.Sprint(s...)
	Print(str, "", "")
}

func PrintColTwo(s ...interface{}) {
	str := fmt.Sprint(s...)
	Print("", str, "")
}

func PrintColThree(s ...interface{}) {
	str := fmt.Sprint(s...)
	Print("", "", str)
}

// print the time in a simple format
func PrintNow() string {
	return time.Now().Format(" sec: 05")
}

// reports the size of the requests queue and burst buffer, and badly pad it to the right of the terminal
func PrintChans(queue chan int, buffer chan bool) {
	for {
		PrintColThree(fmt.Sprintf("request queue: %v, buffer: %v\n", len(queue), len(buffer)))
		time.Sleep(time.Millisecond * 500)
	}
}
