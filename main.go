package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		fmt.Printf("hour: %d, min: %d, sec %d\n", hour, min, sec)

		for line := range clock[0] {

			// For Alarming Every 10 second
			//if sec%10 == 0 {
			//	clock = alarm
			//}

			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
