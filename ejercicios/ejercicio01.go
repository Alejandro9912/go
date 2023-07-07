package ejercicios

import (
	"strconv"
)

func IntString(texto string) (int, string) {
	myInt, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Hubo un error"
	}

	if myInt > 100 {
		return myInt, "Es mayor a 100"
	} else {
		return myInt, "Es menor a 100"
	}
}
