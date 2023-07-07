package ejercicios

import (
	"reflect"
	"strconv"
)

func IntString (texto string) (int,string){
	myInt, err := strconv.Atoi(texto)
	reflect.TypeOf(err)
	var myString string
	if myInt > 100{
		myString = "Es mayor a 100"
	}else{
		myString = "Es menor a 100"
	}
	return myInt, myString
}