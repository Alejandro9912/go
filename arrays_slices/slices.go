package arrays_slices

import "fmt"

//Slice es un Array dinamico

var tablaS []int = []int{1, 2, 3}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func ShowSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice desde un array, desde la posicion 3
	porcion2 := arreglo[:5]  //Slice desde un array, desde la posicion 0 a la 5
	porcion3 := arreglo[6:7] //Slice desde un array, desde la posicion 6 a la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
