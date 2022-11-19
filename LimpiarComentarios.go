package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	result := ""

	centinela := 0
	preCentinela := 0
	centinelaUnaLinea := 0
	controlSalto := 0
	//llamar el archivo plano
	programaFuente, err := ioutil.ReadFile("codigo.sql")
	if err != nil {
		fmt.Println("Hubo un error en la carga del archivo fuente")
	} else {
		filas := tokenizador(programaFuente, "\n")
		for _, line := range filas {
			linea := strings.TrimSpace(line)
			if linea != "" {
				for _, char := range line {
					caracter := string(char)
					if centinela == 0 && centinelaUnaLinea <= 1 {
						if preCentinela == 1 {
							if caracter == "*" {
								centinela = 1
							} else {
								result = result + "/" + caracter
							}
							preCentinela = 0
						} else if centinelaUnaLinea == 1 {
							if caracter == "-" {
								centinelaUnaLinea = 2
								controlSalto = 1
							} else {
								result = result + "-" + caracter
								centinelaUnaLinea = 0
							}
						} else {
							if caracter == "/" {
								preCentinela = 1
							} else if caracter == "-" {
								centinelaUnaLinea = 1
							} else {
								result = result + caracter
							}
						}
					} else if centinela == 1 {
						if preCentinela == 1 {
							if caracter == "/" {
								centinela = 0
								controlSalto = 1
							}
							preCentinela = 0
						} else if caracter == "*" {
							preCentinela = 1
						}
					}
				}
				if centinela == 0 && controlSalto == 0 {
					result = result + "\n"
				}
				centinelaUnaLinea = 0
				controlSalto = 0
			}
		}
	}

	fmt.Printf(result)
}

func tokenizador(cadena []byte, delimitador string) []string {
	texto := string(cadena)
	tokes := strings.Split(texto, delimitador)
	return tokes
}
