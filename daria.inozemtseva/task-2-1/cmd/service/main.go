package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	if n < 1 || n > 1000 {
		return
	}

	for range n {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		if k < 1 || k > 1000 {
			return
		}

		minTemp := 15
		maxTemp := 30

		for range k {
			var symbol string
			var temp int
			if _, err := fmt.Scan(&symbol, &temp); err != nil {
				return
			}

			if symbol == ">=" {
				if temp > minTemp {
					minTemp = temp
				}
			} else if symbol == "<=" {
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
