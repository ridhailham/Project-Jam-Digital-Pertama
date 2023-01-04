package main

import (
	"fmt"
	"time"
)

type placeholder [5]string

func main() {

	zero := placeholder{
		"🀫🀫🀫",
		"🀫  🀫",
		"🀫  🀫",
		"🀫  🀫",
		"🀫🀫🀫",
	}

	one := placeholder{
		"🀫🀫🀫",
		"  🀫 ",
		"  🀫 ",
		"  🀫 ",
		"🀫🀫🀫",
	}

	two := placeholder{
		"🀫🀫🀫",
		"  🀫",
		"🀫🀫🀫",
		"🀫   ",
		"🀫🀫🀫",
	}

	three := placeholder{
		"🀫🀫🀫",
		"  🀫 ",
		"🀫🀫🀫",
		"  🀫 ",
		"🀫🀫🀫",
	}

	four := placeholder{
		"🀫 🀫",
		"🀫 🀫",
		"🀫🀫🀫",
		"  🀫",
		"  🀫",
	}

	five := placeholder{
		"🀫🀫🀫",
		"  🀫 ",
		"🀫🀫🀫",
		"🀫   ",
		"🀫🀫🀫",
	}

	six := placeholder{
		"🀫🀫🀫",
		"🀫   ",
		"🀫🀫🀫",
		"🀫 🀫",
		"🀫🀫🀫",
	}

	seven := placeholder{
		"🀫🀫🀫",
		"   🀫",
		"   🀫",
		"   🀫",
		"   🀫",
	}

	eight := placeholder{
		"🀫🀫🀫",
		"🀫 🀫",
		"🀫🀫🀫",
		"🀫 🀫",
		"🀫🀫🀫",
	}

	nine := placeholder{
		"🀫🀫🀫",
		"🀫  🀫 ",
		"🀫🀫🀫",
		"  🀫 ",
		"🀫🀫🀫",
	}

	colon := [...]string{
		"    ",
		"🀫  🀫",
		"    ",
		"🀫  🀫",
		"    ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	// model for loop lain

	// for _, digit := range digits {
	// 	for _, line := range digit {
	// 		fmt.Println(line)
	// 	}
	// 	fmt.Println()
	// }

	for {
		screen.MoveTopLeft()

		now := time.Now()

		hour, minute, second := now.Hour(), now.Minute(), now.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, minute, second)
		fmt.Println()
		fmt.Println()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				hasil := clock[index][line]
				if digit == colon && second%2 == 0 {
					hasil = "   "
				}
				fmt.Print(hasil, "    ")

			}

		}

		time.Sleep(time.Second)
	}

}
