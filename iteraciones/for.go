package iteraciones

import "fmt"

func Iterar() {
	//continue rompe la ejecucion actual y sigue
	for i := 10; i >= 1; i-= 1 {
		if i==6{
			continue
		}
		fmt.Println(i)
	}
}
