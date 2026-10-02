package main

import "fmt"

func main() {

	var number1, number2 int
	var operator string

	fmt.Print("Enter the first number : ")
	if _, err := fmt.Scanln(&number1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	fmt.Print("Enter the second number : ")
	if _, err := fmt.Scanln(&number2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Print("Enter the operator (+ - * / ) : ")
	_, err := fmt.Scanln(&operator)
	if err != nil || (operator != "+" && operator != "-" && operator != "*" && operator != "/") {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero")
		}else{
			fmt.Println(number1 / number2)
		}
	}
}
