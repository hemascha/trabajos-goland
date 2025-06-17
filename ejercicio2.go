package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {


	reader := bufio.NewReader(os.Stdin)


	fmt.Print("Ingresa la cantidad en bolivares: ")
	bolivaresStr, _ := reader.ReadString('\n')
	bolivaresStr = strings.TrimSpace(bolivaresStr)
	
	bolivares, err := strconv.ParseFloat(bolivaresStr, 64)
	if err != nil {
		fmt.Println("¡Error! Debes ingresar un número")
		return
	}

	fmt.Print("Ingresa el valor de 1 dólar en bolivares: ")
	dolarStr, _ := reader.ReadString('\n')
	dolarStr = strings.TrimSpace(dolarStr)
	
	valorDolar, err := strconv.ParseFloat(dolarStr, 64)
	if err != nil {
		fmt.Println("¡Error! Debes ingresar un número")
		return
	}

	dolares := bolivares / valorDolar

	fmt.Printf("\nResultado: bolivares %.2fbs = $%.2f dólares\n", bolivares, dolares)
}