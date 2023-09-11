/*
Letras de una palabra.
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
	1. Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
	2. Luego, imprimí cada una de las letras.
*/

package main

import "fmt"

func main() {
	palabra := "Estupendo"
	length := len(palabra)
	fmt.Println("Cantidad de letras:", length)
}
