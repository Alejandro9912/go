package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numberOne int
var numberTwo int
var leyend string
var err error

func AddNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer numero: ")
	if scanner.Scan() {
		numberOne, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese el segundo numero: ")
	if scanner.Scan() {
		numberTwo, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese la leyenda: ")
	if scanner.Scan() {
		leyend = scanner.Text()
	}
	fmt.Println(leyend, numberOne*numberTwo)
}
