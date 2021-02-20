/*
	Ejemplo 1: Server with Go
	@whary11 Luis Fernando Raga
*/
 
package main
 
import (
	"fmt"// Imprimir en consola
	"io"// Ayuda a escribir en la respuesta
	"log"//Loguear si algo sale mal
	"net/http"// El paquete HTTP
	"encoding/json"
)

type Person map[string]int
 
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, p *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response, err := getJsonResponse();
		if err != nil {
			panic(err)
		}
			fmt.Println(string(response))
			io.WriteString(w, string(response))
	})
	http.HandleFunc("/hola", func(w http.ResponseWriter, p *http.Request) {
		io.WriteString(w, "Solicitaste hola")
	})
	port := ":8080" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Ingresa a localost:"+port)
	fmt.Println("Aquí podrás saber mas de mi")
	log.Fatal(http.ListenAndServe(port, nil))
}

func getJsonResponse()([]byte, error) {
    persona := make(map[string]string)
	persona["name"] = "Luis Fernando"
	persona["surname"] = "Raga Renteria"
	persona["age"] = "31"
	persona["stack"] = "Go, JS, HTML, CSS"

    return json.MarshalIndent(persona, "", "  ")
}