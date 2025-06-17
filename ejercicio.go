package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("¿Cómo te llamas? ")
	scanner.Scan()
	nombre := strings.TrimSpace(scanner.Text())

	fmt.Print("Ingresa tu edad: ")
	scanner.Scan()
	edadStr := scanner.Text()

	edad, err := strconv.Atoi(edadStr)
	if err != nil {
		fmt.Println("¡Error! Debes ingresar un número.")
		return
	}

	meses := edad * 12
	fmt.Printf("%s, has vivido aproximadamente %d meses.\n", nombre, meses)
}