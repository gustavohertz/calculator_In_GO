package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Calculadora em Go")
	fmt.Println("------------------")

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Erro: Divisão por zero!")
			return
		}
	default:
		fmt.Println("Operador inválido!")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
