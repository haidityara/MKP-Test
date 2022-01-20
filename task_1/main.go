package main

import "fmt"

func main() {

	rows := 9
	// define space as mid-rows
	space := rows / 2
	num := 1

	// loop through i <= rows
	for i := 1; i <= rows; i++ {
		// loop to get outer start left
		for j := 1; j <= space; j++ {
			fmt.Printf("*")
		}

		// loop as space of diamond count
		for k := 1; k <= num; k++ {
			// if first row no need draw space
			if i == 1 {
				fmt.Printf("*")
				//  also, for last row
			} else if i == rows {
				fmt.Printf("*")
			} else {
				// if in the middle of rows
				if space == 0 {
					// if first col add star
					if k == 1 {
						fmt.Printf("*")
						// also, for last row add start
					} else if k == rows {
						fmt.Printf("*")
						// else draw -
					} else {
						fmt.Printf("-")
					}
				} else {
					// main role to draw to price space
					fmt.Printf(" ")
				}
			}

		}

		// draw right star
		for j := 1; j <= space; j++ {
			fmt.Printf("*")
		}

		fmt.Printf("\n")

		// set space and num
		if i <= rows/2 {
			space = space - 1
			num = num + 2
		} else {
			space = space + 1
			num = num - 2
		}
	}

}
