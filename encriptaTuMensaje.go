package main

import (
	"strings"
)

// encriptar encripta el mensaje usando la clave proporcionada.
func encriptar(clave, mensaje string) string {
	// Si el mensaje es nulo o vacío, o la clave está vacía, retornamos una cadena vacía.
	if mensaje == "" {
		return ""
	}

	// Si la clave está vacía, usamos el valor por defecto "DCJ".
	if clave == "" {
		clave = "DCJ"
	}

	var resultado strings.Builder

	// Defino las vocales
	vocales := "aeiouAEIOU"

	for _, char := range mensaje {
		// Si el carácter es una vocal, agregar la clave antes de él
		if strings.ContainsRune(vocales, char) {
			resultado.WriteString(clave)
		}
		// Agregar el carácter al resultado
		resultado.WriteRune(char)
	}

	return resultado.String()
}

func main() {
	clave := "dcj"
	mensaje := "I love prOgrAmming!"
	resultado := encriptar(clave, mensaje)
	println(resultado) //se observa por consola: dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!
}

//Si no queremos definir las vocales podemos importar unicode ya que es otra de las formas de realizar el ejercicio.
