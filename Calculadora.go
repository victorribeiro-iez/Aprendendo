package main

import (
	"fmt"
	"os"
)

const somar = 1
const subtrair = 2
const dividir = 3
const multiplicacar = 4

func soma(n1 float32, n2 float32) float32 {
	return n1 + n2
}

func subtracao(n1 float32, n2 float32) float32 {
	return n1 - n2
}

func multiplicacao(n1 float32, n2 float32) float32 {
	return n1 * n2
}

func divisao(n1 float32, n2 float32) float32 {
	return n1 / n2
}

func main() {
	fmt.Println("----------Calculadora-----------")

	var operacacao string
	var valor1 float32
	var valor2 float32
	var resultado float32

	fmt.Println("Escolha qual operação deseja fazer")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("0 - Sair")
	fmt.Println("------------------------------")
	fmt.Scan(&operacacao)
	fmt.Println("Informe o primeiro valor")
	fmt.Scan(&valor1)
	fmt.Println("Informe o segundo valor")
	fmt.Scan(&valor2)
	switch operacacao {
	case "0":
		{
			os.Exit(-1)
		}
	case "1":
		{
			resultado = soma(valor1, valor2)
			fmt.Println("O resultado da soma é : ", resultado)
		}
	case "2":
		{
			resultado = subtracao(valor1, valor2)
			fmt.Println("O resultado da subtração é : ", resultado)
		}
	case "3":
		{
			resultado = multiplicacao(valor1, valor2)
			fmt.Println("O resultado da multiplicação é : ", resultado)
		}
	case "4":
		{
			resultado = divisao(valor1, valor2)
			fmt.Println("O resultado da divisão é : ", resultado)
		}
	default:
		{
			fmt.Println("Operação inválida")
		}

	}

}
