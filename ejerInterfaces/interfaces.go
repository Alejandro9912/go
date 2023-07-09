package ejerinterfaces

import (
	"fmt"

	"github.com/Alejandro9912/go/interfaces"
)

func HumanosRespirando(hu interfaces.Humano){
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}