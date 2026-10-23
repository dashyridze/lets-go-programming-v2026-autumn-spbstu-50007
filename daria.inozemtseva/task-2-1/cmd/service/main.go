package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)

	if N < 1 || N > 1000 {
		return
	}

	for i := 0; i < N; i++ {
		var K int
		fmt.Scan(&K)

		if K < 1 || K > 1000 {
			return
		}

		var min int = 15
		var max int = 30

		for j := 0; j < K; j++ {
			var symbol string
			var temp int
			fmt.Scan(&symbol, &temp)

			if symbol == ">=" {
				if temp > min {
					min = temp
				}
			} else if symbol == "<=" {
				if temp < max {
					max = temp
				}
			}

			if min <= max {
				fmt.Println(min)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
