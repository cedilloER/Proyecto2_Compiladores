package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	sinComentarios := ""
	programaFuente, err := ioutil.ReadFile("codigo.sql")
	if err != nil {
		fmt.Println("Hubo un error en la carga del archivo fuente")
	} else {
		sinComentarios, err = deleteComentarios(programaFuente)
		if err != nil {
			fmt.Println("Se encontraron errores")
		} else {
			fmt.Printf(sinComentarios)
		}
	}

	prueba := [7]string{"create", "_hola", "23.73", "int", ")", "DataBase", `"mi mundo cruel"`}

	for _, p := range prueba {
		println(getTipo(p))
	}
}

// Esta funcion es la encargada de eliminar los comentario y saltos de linea
func deleteComentarios(cadena []byte) (string, error) {
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
							return limpieza(result), io.EOF
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
		return result, io.EOF
	}

	return result, nil
}

// Funcion que ayuda a limpiar saltos de linea
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

// devuelve el tipo de lexema
func getTipo(lexema string) string {
	rgExpre := getExpRegulares()
	for _, exp := range rgExpre {
		regex_ruta := regexp.MustCompile(exp[1])
		if regex_ruta.FindStringSubmatch(lexema) != nil {
			return exp[0]
		}
	}
	return "Sin valor"
}

// Funcion encargada de agregar los tokens a la tabla de simbolos
func agregarTokens(tokens []string, tablaSimbolos map[int]map[string]string) map[int]map[string]string {
	for _, lexema := range tokens {
		tipo := getTipo(lexema)
		if tipo != "" {
			aux := make(map[string]string)
			aux["lexema"] = lexema
			aux["tipo"] = tipo
			tablaSimbolos[len(tablaSimbolos)] = aux
		} else {
			println("El lexema no cumple con ningun tipo")
		}
	}
	return tablaSimbolos
}

// Funcion encargada de obtener las expresiones regulares y convertirlas en un map
func getExpRegulares() [][]string {
	ExpRegulares := [][]string{
		{"create", `^[cC][rR][eE][aA][tT][eE]$`},
		{"use", `^[uU][sS][eE]$`},
		{"database", `^[dD][aA][tT][aA][bB][aA][sS][eE]$`},
		{"table", `^[tT][aA][bB][lL][eE]$`},
		{"int", `^[iI][nN][tT]$`},
		{"varchar", `^[vV][aA][rR][cC][hH][aA][rR]$`},
		{"float", `^[fF][lL][oO][aA][tT]$`},
		{"date", `^[dD][aA][tT][eE]$`},
		{"char", `^[cC][hH][aA][rR]$`},
		{"boolean", `^[bB][oO][oO][lL][eE][aA][nN]$`},
		{"insert", `^[iI][nN][sS][eE][rR][tT]$`},
		{"into", `^[iI][nN][tT][oO]$`},
		{"values", `^[vV][aA][lL][uU][eE][sS]$`},
		{"select", `^[sS][eE][lL][eE][cC][tT]$`},
		{"from", `^[fF][rR][oO][mM]$`},
		{"delete", `^[dD][eE][lL][eE][tT][eE]$`},
		{"where", `^[wW][hH][eE][rR][eE]$`},
		{"literal", `^"[[:print:]]+"$`},
		{"numero", `^[0-9]+(.[0-9]+)?$`},
		{"comparacion", `^(=|<|>|<=|>=|<>|!=|!<|!>)$`},
		{"delimitador", `^[;]$`},
		{"separador", `^[,]$`},
		{"parantesisAbierto", `^[(]$`},
		{"parentesisCerrado", `^[)]$`},
		{"id", `^([0-9_]+)?[[:alpha:]]([[:alnum:]_]+)?$`}}
	return ExpRegulares
}

// Inicia la tabla de simbolos con los valores iniciales de la sintaxis del lenguaje
func getTablaDeSimbolos() map[int]map[string]string {
	tablaDeSimbolos := make(map[int]map[string]string)
	valoresIniciales := getExpRegulares()

	for _, vi := range valoresIniciales {
		v := make(map[string]string)
		v["tipo"] = vi[0]
		v["lexema"] = ""
		tablaDeSimbolos[len(tablaDeSimbolos)] = v
	}

	return tablaDeSimbolos
}

// Funcion que separa una cadena segun el delimitador que se asigne
func tokenizador(cadena []byte, delimitador string) []string {
	texto := string(cadena)
	tokes := strings.Split(texto, delimitador)
	return tokes
}
