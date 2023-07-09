package deferPanic

import "fmt"

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje Final")
	fmt.Println("Este es el segundo mensaje")
}
