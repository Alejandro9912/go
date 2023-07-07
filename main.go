package main

import (
	"fmt"
	"github.com/Alejandro9912/go/variables"
)

func main() {
	/* fmt.Println("Hola Mundo")
	variables.MuestroEnteros()
	variables.RestoVariables() */
	estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
