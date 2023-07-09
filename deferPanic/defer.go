package deferPanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	//defer indica cual o cuales son las ejecuciones finales
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje Final")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	//Panic es un aborto en la ejecucion
	//Recovery para recuperarse de un aborto
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
