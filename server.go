/*
	Ejemplo 1: Server with Go
	@whary11 Luis Fernando Raga
*/

package main

import (
	"encoding/json"
	"fmt" // Imprimir en consola
	"io"  // Ayuda a escribir en la respuesta
	"log" //Loguear si algo sale mal
	"net"
	"net/http" // El paquete HTTP
	"strings"
)

type Person map[string]int

func main() {
	fmt.Printf("%d\n", getIvaValue(10000, 10))

	http.HandleFunc("/", func(w http.ResponseWriter, p *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response, err := getJsonResponse()
		if err != nil {
			panic(err)
		}
		// fmt.Println(string(response))
		// fmt.Println(getOutboundIP())
		io.WriteString(w, string(response))
	})
	http.HandleFunc("/hola", func(w http.ResponseWriter, p *http.Request) {
		io.WriteString(w, "Solicitaste hola")
	})
	port := ":4000" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Ingresa a: " + getOutboundIP() + ":" + port)
	fmt.Println("Aquí podrás saber mas de mi")
	log.Fatal(http.ListenAndServe(port, nil))
}

func getJsonResponse() ([]byte, error) {
	persona := make(map[string]string)
	persona["name"] = "Luis Fernando"
	persona["surname"] = "Raga Renteria"
	persona["age"] = "31"
	persona["stack"] = "Go, JS, HTML, CSS"

	return json.MarshalIndent(persona, "", "  ")
}
func getOutboundIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")

	// HandleError("net.Dial: ", err)
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func getIvaValue(value int, iva int) int {
	return ((value * iva) / 100)
}
