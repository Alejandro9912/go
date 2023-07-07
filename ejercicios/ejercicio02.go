package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TenMultip() string {
	scanner := bufio.NewScanner(os.Stdin)
	var texto string

	fmt.Println("Ingrese el numero: ")
	if scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El valor ingresado no es valido " + err.Error())
			TenMultip()
		} else {
			for i := 1; i <= 10; i++ {
				//fmt.Printf("%d * %d = %d \n", number,i,number*i)
				texto += fmt.Sprintln(i, "*", number, "=", number*i)
			}
		}
	}
	fmt.Println(texto)
	return texto
}
