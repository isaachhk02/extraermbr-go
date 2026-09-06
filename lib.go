package main
import (
	"fmt"
	"os"
)

func extraer_tabla(dispositivo string, ruta_binario string) {
	fmt.Println("Abriendo ", dispositivo)
	dev,err := os.Open(dispositivo)
	if err != nil {
		panic("Error al abrir el dispositivo")
	}
	defer dev.Close()

	buffer := make([]byte,512)
	fmt.Println("Leyendo tabla de particiones: ", dispositivo)
	n,err := dev.Read(buffer)
	if err != nil {
		panic("Error al leer la tabla de particiones")
	}
	fmt.Println("Leido ", int(n))
	binario,err := os.Create(ruta_binario)
	if err != nil {
		panic("Error al crear el archivo binario")
	}
	b,err := binario.Write(buffer)
	if err != nil {
		panic("Error al escribir al escribir al archivo binario")
	}
	fmt.Println("El archivo ", ruta_binario, "se ha escrito correctamente")
	fmt.Println("Bytes: ", int(b))
	binario.Close()
}
