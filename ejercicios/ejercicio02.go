package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TenMultip() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el numero: ")
	if scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El valor ingresado no es valido " + err.Error())
			TenMultip()
		} else {
			for i := 1; i <= 10; i++ {
				//fmt.Printf("%d * %d = %d \n", number,i,number*i)
				fmt.Println(i, "*", number, "=", number*i)
			}
		}
	}
}
