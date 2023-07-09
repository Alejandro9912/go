package main

import (
	"github.com/Alejandro9912/go/webserver"
)

func main() {
	/* fmt.Println("Hola Mundo")
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os) // %s define que ingreso un String
	}

	numero, texto := ejercicios.IntString("fffff")
	fmt.Println(numero, reflect.TypeOf(numero))
	fmt.Println(texto)

	teclado.AddNumbers()
	iteraciones.Iterar()
	ejercicios.TenMultip()
	files.RecordTable()
	files.SumTable()
	files.LeoArchivo()
	files.ReadArch()
	funciones.LlamarClosure()
	funciones.Exponencia(2)
	arrays_slices.ShowArray()
	funciones.Calculos()
	arrays_slices.ShowSlice()
	arrays_slices.Capacidad()
	mapas.ShowMap()
	users.AltaUser()

	Pedro := new(models.Hombre)
	ejerinterfaces.HumanosRespirando(Pedro)

	Maria := new(models.Mujer)
	ejerinterfaces.HumanosRespirando(Maria)
	deferPanic.VemosDefer()
	deferPanic.EjemploPanic()
	canal1 := make(chan bool)
	go goroutines.MiNombreLento("Juan Calderon", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy Aqui")*/
	webserver.MiWebServer()
}
