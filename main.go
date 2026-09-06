// Creado por isaachhk02 al dia 22/08/2026

package main

import (
	"fmt"
	"os"
)

func main() {

	argumentos := os.Args

	if len(argumentos) < 3 {
		fmt.Println("mbrextractor-go [dispositivo] [archivo_salida]")
		fmt.Println("dispositivo: Ruta del dispositivo de bloques (ejemplo: /dev/sda) para extraer la tabla de particiones")
		fmt.Println("archivo_salida: Ruta donde se guardara la tabla de particiones en un archivo binario")
	} else {
		fmt.Println("Extrayendo tabla de particiones: ", argumentos[1])
		extraer_tabla(argumentos[1],argumentos[2])
	}
}
