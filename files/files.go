package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Alejandro9912/go/ejercicios"
)

var fileName string = "./files/txt/tableM.txt"

func RecordTable() {
	var texto string = ejercicios.TenMultip()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el acrhivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumTable() {
	var texto string = ejercicios.TenMultip()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error en el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error en el WriteString " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
	}
	fmt.Println(string(archivo))
}

func ReadArch(){
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan(){
		registro := scanner.Text()
		fmt.Println("> "+ registro)
	}
	archivo.Close()
}
