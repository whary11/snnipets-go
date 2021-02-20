package main
import "./db/leer/obtenerContactos"
import "fmt"
func main() {
	contactos, err := obtenerContactos()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}
	for _, contacto := range contactos {
		fmt.Printf("%v\n", contacto)
	}
}