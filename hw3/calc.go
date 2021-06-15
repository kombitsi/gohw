package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
getOperation:
	var operation string
	fmt.Print("Выберите действие: +, -, *, /, x^y, max, min: ")
	fmt.Scanln(&operation)
	var num1 string
	fmt.Print("Введите первое число:")
	fmt.Scanln(&num1)
	if _, err := strconv.Atoi(num1); err == nil {
		fmt.Printf("%q Похоже на число.\n", num1)
	} else {
		fmt.Println("Это не число")
	}

	var num2 string
	fmt.Print("Ввидите второе число:")
	fmt.Scanln(&num2)
	if _, err := strconv.Atoi(num2); err == nil {
		fmt.Printf("%q Похоже на число.\n", num2)

	} else {
		fmt.Println("Это не число")
	}

	switch operation {
	case "+":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(Add(firstNumber, secondNumber))
	case "-":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(Subtract(firstNumber, secondNumber))
	case "*":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(Multiply(firstNumber, secondNumber))
	case "/":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(Divide(firstNumber, secondNumber))
	case "pow":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(Pow(firstNumber, secondNumber))
	case "max":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(Pow(firstNumber, secondNumber))
	case "min":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(Pow(firstNumber, secondNumber))
	default:
		fmt.Println("Invalid operation selected. Please try again!")
		goto getOperation
	}
}
func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}

func Add(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}
func Divide(num1, num2 float64) float64 {
	return num1 / num2
}

func Pow(num1, num2 float64) float64 {
	p := math.Pow(num1, num2)
	return p
}
func Max(num1, num2 float64) float64 {
	max := math.Max(num1, num2)
	return max
}

func Min(num1, num2 float64) float64 {
	min := math.Min(num1, num2)
	return min
}
