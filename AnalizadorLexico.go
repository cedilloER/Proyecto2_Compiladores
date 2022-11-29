package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	sinComentarios := ""
	programaFuente, err := ioutil.ReadFile("codigo.sql")
	if err != nil {
		fmt.Println("Hubo un error en la carga del archivo fuente")
	} else {
		sinComentarios = deleteComentarios(programaFuente)
	}
	fmt.Println("\n\n\n _________________________________________________________")
	fmt.Printf(sinComentarios)
}

func deleteComentarios(cadena []byte) string {
	result := ""
	centinela := 0
	preCentinela := 0
	preError := 0
	centinelaUnaLinea := 0

	linePosibleError := 0

	filas := tokenizador(cadena, "\n")
	for i, line := range filas {
		linea := strings.TrimSpace(line)
		if linea != "" {
			for _, char := range line {
				caracter := string(char)
				if centinela == 0 && centinelaUnaLinea <= 1 {
					if preCentinela == 1 {
						if caracter == "*" {
							centinela = 1
							linePosibleError = i + 1
						} else {
							result = result + "/" + caracter
						}
						preCentinela = 0
					} else if centinelaUnaLinea == 1 {
						if caracter == "-" {
							centinelaUnaLinea = 2
						} else {
							result = result + "-" + caracter
							centinelaUnaLinea = 0
						}
					} else if preError == 1 {
						if caracter == "/" {
							fmt.Printf("Error en linea ")
							fmt.Println(i + 1)
						} else {
							result = result + "*" + caracter
							preError = 0
						}
					} else {
						if caracter == "/" {
							preCentinela = 1
						} else if caracter == "-" {
							centinelaUnaLinea = 1
						} else if caracter == "*" {
							preError = 1
						} else {
							result = result + caracter
						}
					}
				} else if centinela == 1 {
					if preCentinela == 1 {
						if caracter == "/" {
							centinela = 0
							linePosibleError = 0
						}
						preCentinela = 0
					} else if caracter == "*" {
						preCentinela = 1
					}
				}
			}

			result = result + "\n"

			centinelaUnaLinea = 0
		}
	}

	result = limpieza(result)

	if linePosibleError != 0 {
		fmt.Printf("Error en la linea ")
		fmt.Println(linePosibleError)
	}

	return result
}

func limpieza(cadena string) string {
	resultLimpio := ""

	filas := tokenizador([]byte(cadena), "\n")

	for _, linea := range filas {
		line := strings.TrimSpace(linea)
		if line != "" {
			resultLimpio = resultLimpio + line + "\n"
		}
	}

	return resultLimpio
}

func tokenizador(cadena []byte, delimitador string) []string {
	texto := string(cadena)
	tokes := strings.Split(texto, delimitador)
	return tokes
}
