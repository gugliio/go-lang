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
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])

	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}

	operador2, err2 := strconv.Atoi(valores[1])
	if err2 != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}

	fmt.Println(operador1 + operador2)
}
