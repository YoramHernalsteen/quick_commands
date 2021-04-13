package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type timeFormat struct {
	t int
	h int
	m int
	s int
}

func main() {
	fmt.Printf("\n")
	if minutes, err := strconv.Atoi(os.Args[1:][0]); err == nil {
		endTime := time.Now().Local().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(minutes) +
			time.Second*time.Duration(0))

		for range time.Tick(1 * time.Second) {
			timeRemaining := getTimeRemaining(endTime)

			if timeRemaining.t <= 0 {
				fmt.Printf("\033[2K\r")
				fmt.Println("\nRingRingRing, Timer's up!")
				break
			}

			if timeRemaining.h == 0 && timeRemaining.m == 0 {
				fmt.Printf("\033[2K\rRemaining: %d", timeRemaining.s)
			} else if timeRemaining.h == 0 {
				fmt.Printf("\033[2K\rRemaining: %d:%d", timeRemaining.m, timeRemaining.s)
			} else {
				fmt.Printf("\033[2K\rRemaining: %d:%d:%d", timeRemaining.h, timeRemaining.m, timeRemaining.s)
			}

		}
	} else {
		fmt.Println("Please write an integer.")
		os.Exit(1)
	}
}

func getTimeRemaining(t time.Time) timeFormat {
	current := time.Now()
	difference := t.Sub(current)

	total := int(difference.Seconds())
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return timeFormat{
		t: total,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
